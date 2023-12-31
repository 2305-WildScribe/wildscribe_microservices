package model

import "wildscribe.com/gen"

// AdventureToProto converts a Adventure struct into a
// generated proto counterpart.
func AdventureToProto(adventure *Adventure) *gen.Adventure {
	return &gen.Adventure{
		AdventureId:        adventure.Adventure_id,
		UserId:             adventure.User_id,
		Activity:           adventure.Activity,
		Date:               adventure.Date,
		ImageUrl:           adventure.Image_url,
		StressLevel:        adventure.Stress_level,
		HoursSlept:         adventure.Hours_slept,
		SleepStressNotes:   adventure.Sleep_stress_notes,
		Hydration:          adventure.Hydration,
		Diet:               adventure.Diet,
		DietHydrationNotes: adventure.Diet_hydration_notes,
		BetaNotes:          adventure.Beta_notes,
		Lat:                adventure.Lat,
		Lon:                adventure.Lon,
	}
}

// AdventureFromProto converts a generated proto counterpart
// into a Adventure struct.
func AdventureFromProto(m *gen.Adventure) *Adventure {
	return &Adventure{
		Adventure_id:         m.AdventureId,
		User_id:              m.UserId,
		Activity:             m.Activity,
		Date:                 m.Date,
		Image_url:            m.ImageUrl,
		Stress_level:         m.StressLevel,
		Hours_slept:          m.HoursSlept,
		Sleep_stress_notes:   m.SleepStressNotes,
		Hydration:            m.Hydration,
		Diet:                 m.Diet,
		Diet_hydration_notes: m.DietHydrationNotes,
		Beta_notes:           m.BetaNotes,
		Lat:                  m.Lat,
		Lon:                  m.Lon,
	}
}

// AdventureSliceFromProto converts a slice of Adventure protos to a slice of model.Adventure.
func AdventureSliceFromProto(protoAdventures []*gen.Adventure) []*Adventure {
	var adventures []*Adventure
	for _, protoAdventure := range protoAdventures {
		adventures = append(adventures, AdventureFromProto(protoAdventure))
	}
	return adventures
}

// AdventureSliceToProto converts a slice of model.Adventure to a slice of gen.Adventure.
func AdventureSliceToProto(adventures []*Adventure) []*gen.Adventure {
	var protoAdventures []*gen.Adventure
	for _, adventure := range adventures {
		protoAdventures = append(protoAdventures, AdventureToProto(adventure))
	}
	return protoAdventures
}
