package model

import (
	"wildscribe.com/gen"
)

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
	}
}
