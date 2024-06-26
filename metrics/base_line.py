from locust import HttpUser, task, between
import random
import logging
# from faker import Faker



class User(HttpUser):
    host = "http://localhost:8080/api/v0"
    wait_time = between(1, 5)
    weight = 1

    def on_start(self):
        self.login_index = random.randint(1, 499)
        self.user_id = None
        self.adventure_id = None

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
            if response.elapsed.total_seconds() > 1.0:
                response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
            elif response.status_code == 200:
                self.user_id = response.json()["data"]["attributes"]["user_id"]
                response.success()
            else: 
                response.failure(f"Request failed : {response.json()}")
   
    @task
    def GetUserAdventures(self):
        if self.user_id is None:
            self.LoginUser()

        endpoint = "/user/adventures"

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
            if response.elapsed.total_seconds() > 1.0:
                response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
            elif response.status_code != 200:
                response.failure(f"Request has diffrent status code: {response.status_code}")
            elif response.status_code == 200:
                response.success()
            else: 
                response.failure(f"Request failed : {response.json()}")
   
    # @task(5)
    # def CreateAdventure(self):
    #     if self.user_id is None:
    #         self.LoginUser()

    #     endpoint = "/adventure"

    #     payload = { 
    #                 "data":{
    #                     "type": "adventures",
    #                     "attributes": {
    #                         "user_id": self.user_id,
    #                         "activity": "Running",
    #                         "date": "10/11/2023"
    #                     }
    #                  }
    #             }
    #     headers = { "Content-Type": "application/json "}

    #     with self.client.post(endpoint, json=payload, headers=headers, catch_response=True) as response:
    #         if response.elapsed.total_seconds() > 1.0:
    #             response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
    #         elif response.status_code == 201:
    #             self.adventure_id = response.json()["data"]["attributes"]["adventure_id"]
    #             response.success()
    #         elif response.status_code != 200:
    #             response.failure(f"Request has diffrent status code: {response.status_code}")
    #         else: 
    #             response.failure(f"Request failed : {response.json()}")
   
  
    # @task(2)
    # def UpdateAdventure(self):

    #     if self.user_id is None:
    #         return
    #     if self.adventure_id is None:
    #         return
        
    #     endpoint = "/adventure"

    #     payload = { 
    #                 "data":{
    #                     "type": "adventures",
    #                     "attributes": {
    #                         "adventure_id": self.adventure_id,
    #                         "user_id": self.user_id,
    #                         "activity": "Walking",
    #                         "date": "10/11/2023"
    #                     }
    #                  }
    #             }
    #     headers = { "Content-Type": "application/json "}

    #     with self.client.put(endpoint, json=payload, headers=headers, catch_response=True) as response:
    #         if response.elapsed.total_seconds() > 1.0:
    #             response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
    #         elif response.status_code == 200:
    #             response.success()
    #         elif response.status_code == 500:
    #             response.success()
    #         else: 
    #             response.failure(f"Request failed : {response.text}")
    # @task(2)
    # def DeleteAdventure(self):
    #     if self.adventure_id is None:
    #         return

    #     endpoint = "/adventure"

    #     payload = { 
    #                 "data":{
    #                     "type": "adventures",
    #                     "attributes": {
    #                         "adventure_id": self.adventure_id
    #                     }
    #                  }
    #             }
    #     headers = { "Content-Type": "application/json "}

    #     with self.client.delete(endpoint, json=payload, headers=headers, catch_response=True) as response:
    #         if response.elapsed.total_seconds() > 1.0:
    #             response.failure(f"Request took to long: {response.elapsed.total_seconds()}")
    #         elif response.status_code == 200:
    #             self.adventure_id = None
    #             response.success()
    #         elif response.status_code != 200:
    #             response.failure(f"Request has diffrent status code: {response.status_code}")
    #         else: 
    #             response.failure(f"Request failed : {response.json()}")

    # @task(1)
    # def Logout(self):
    #     self.user_id = None
   