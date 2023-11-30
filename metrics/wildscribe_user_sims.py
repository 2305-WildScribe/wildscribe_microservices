from locust import HttpUser, task, between
import random
# from faker import Faker

class User(HttpUser):
    host = "http://localhost:8080/api/v0"
    wait_time = between(1, 5)
    weight = 1

    def on_start(self):
        self.login_index = random.randint(1, 4999)
        self.user_id = None
        self.LoginUser()

    @task
    def LoginUser(self):
        if self.user_id is not None:
            return

        endpoint = "/user"

        payload = { 
                    "data":{
                        "type": "user",
                        "attributes": {
                            "email": f"me{self.login_index}@gmail.com",
                            "password": f"Password{self.login_index}"
                        }
                     }
                }
        headers = { "Content-Type": "application/json "}

        with self.client.post(endpoint, json=payload, headers=headers, catch_response=True) as response:
            self.user_id = response.json()["data"]["attributes"]["user_id"]
            if response.status_code == 200:
                response.success()
            elif response.status_code != 200:
                response.failure(f"Request has diffrent status code: {response.status_code}")
            elif response.elapsed.total_seconds() > 1.0:
                response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
            else: 
                response.failure(f"Request failed with status: {response.reason}")
   
    @task
    def GetUserAdventures(self):
        endpoint = "/adventures"

        payload = { 
                    "data":{
                        "type": "adventures",
                        "attributes": {
                            "user_id": self.user_id,
                        }
                     }
                }
        headers = { "Content-Type": "application/json "}

        with self.client.post(endpoint, json=payload, headers=headers, catch_response=True) as response:
            if response.status_code == 200:
                response.success()
            elif response.status_code != 200:
                response.failure(f"Request has diffrent status code: {response.status_code}")
            elif response.elapsed.total_seconds() > 1.0:
                response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
            else: 
                response.failure(f"Request failed with status: {response.reason}")
   
#     @task(4)
#     def GetAllTeas(self):

#         endpoint = "/teas"

#         with self.client.get(endpoint, catch_response=True) as response:
#             if response.elapsed.total_seconds() > 1.0:
#                 response.failure(f"Request took too long: {response.elapsed.total_seconds()} seconds")
#             elif response.status_code == 200:
#                 response.success()
#             else: 
#                 response.failure("Request Failed")

#     @task(3)
#     def GetATea(self):

#         tea_id = random.randint(1, 10000)

#         endpoint = f"/teas/{tea_id}"

#         with self.client.get(endpoint, catch_response=True) as response:
#             response_id = response.json()["data"]["id"]
#             if response.elapsed.total_seconds() > 1.0:
#                 response.failure(f"Request took too long: {response.elapsed.total_seconds()} seconds")
#             elif response.status_code == 200 and tea_id == int(response_id):
#                 response.success()
#             elif response.status_code == 410:
#                 response.success()
#             else: 
#                 response.failure(f"Request failed with status: {response.reason}")

#     @task
#     def GetACustomer(self):
#         endpoint = "/customer"

#         payload = {
#             "data":{
#                 "type": "customer",
#                 "attributes": {
#                     "id": self.customer_id
#                 }
#             }
#         }

#         headers = { "Content-Type": "application/json "}

#         with self.client.post(endpoint, json=payload, headers=headers, catch_response=True) as response:
#             response_id = response.json()["data"]["id"]
#             if response.status_code == 200 and self.customer_id == int(response_id):
#                 response.success()
#             elif response.elapsed.total_seconds() < 1.0:
#                 response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
#             else: 
#                 response.failure(f"Request failed with status: {response.reason}")

#     @task(2)
#     def Subscribe(self):
#         endpoint = "/subscribe"

#         if self.set_tea_id == None:
#             self.set_tea_id = random.randint(1,10000)

#         status = ["pending", "active", "inactive"]
#         payload = {
#             "data": {
#                 "type": "subscription",
#                 "attributes": {
#                     "customer_id": self.customer_id,
#                     "tea_id": self.set_tea_id,
#                     "status": status[random.randint(0,2)],
#                     "title": "Earl Grey",
#                     "price": self.faker.pricetag(),
#                     "frequency": "Monthly"
#                 }
#             }
#         }

#         headers = {
#             "Content-Type": "application/json"
#         }

#         with self.client.post(endpoint, json=payload, headers=headers, catch_response=True) as response:
#             if response.status_code == 200:
#                 response.success()
#             elif response.elapsed.total_seconds() > 1.0:
#                 response.failure(f"Request took too long: {response.elapsed.total_seconds()} seconds")


#     @task(2)
#     def Unsubcribe(self):
#         endpoint = "/unsubscribe"
        
#         if self.set_tea_id == None:
#             self.Subscribe()
        
#         payload = {
#             "data": {
#                 "type": "subscription",
#                 "attributes": {
#                     "customer_id": self.customer_id,
#                     "tea_id": self.set_tea_id
#                 }
#             }
#         }

#         headers = {
#             "Content-Type": "application/json"
#         }

#         with self.client.post(endpoint, json=payload, headers=headers, catch_response=True) as response:
#             if response.status_code == 200:
#                 self.set_tea_id = None
#                 response.success()
#             elif response.elapsed.total_seconds() > 1.0:
#                 response.failure(f"Request took too long: {response.elapsed.total_seconds()} seconds")


# class LurkerUser(HttpUser):
#     host = "http://localhost:3000/api/v0"
#     wait_time = between(1, 5)
#     weight = 3

#     @task(2)
#     def GetAllTeas(self):

#         endpoint = "/teas"

#         with self.client.get(endpoint, catch_response=True) as response:
#             if response.elapsed.total_seconds() > 1.0:
#                 response.failure(f"Request took too long: {response.elapsed.total_seconds()} seconds")
#             elif response.status_code == 200:
#                 response.success()
#             else: 
#                 response.failure("Request Failed")

#     @task
#     def GetATea(self):

#         tea_id = random.randint(1, 10000)

#         endpoint = f"/teas/{tea_id}"

#         with self.client.get(endpoint, catch_response=True) as response:
#             response_id = response.json()["data"]["id"]
#             if response.elapsed.total_seconds() > 1.0:
#                 response.failure(f"Request took too long: {response.elapsed.total_seconds()} seconds")
#             elif response.status_code == 200 and tea_id == int(response_id):
#                 response.success()
#             elif response.status_code == 410:
#                 response.success()
#             else: 
#                 response.failure(f"Request failed with status: {response.reason}")