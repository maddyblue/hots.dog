package main

import "testing"

func TestXML(t *testing.T) {
	tests := map[string]string{
		"Abil,AlarakTelekinesis,Range":                                                              "8",
		"Behavior,AnubarakAcidDrenchedMandibles,Modification.DamageDealtFraction[Basic]":            "0.7",
		"Effect,AlarakLightningSurgeExtendedLightningTalent5TokenModifyPlayer,EffectArray[0].Value": "1.2",
		"Effect,AlarakLightningSurgeHeroDamage,MultiplicativeModifierArray[Base].Modifier":          "1",
		"Effect,AlarakRuthlessMomentumModifyCooldowns,Cost[0].CooldownTimeUse":                      "-0.0625",
		"Effect,AlarakSadismAdd6TokensBladeoftheHighlordTalent,Value":                               "7",
		"Effect,AnaOverdoseAttackPoisonTraitDotDamageTokenInitialSingleDoseApplyBehavior,Value":     "1",
		"Effect,RexxarAnimalHusbandryModifyTokenCounter,Value":                                      "1",
		"Talent,AlarakAppliedForce,AbilityModificationArray[1].Modifications[0].Value":              "1.600000",
		"Unit,HeroAna,Sight":                                                                        "12",
		"Value":                                                                                     "0",
	}
	var x XML
	for _, p := range []string{
		"mods/heroesdata.stormmod/base.stormdata/GameData/EffectData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/Heroes/AnubarakData/AnubarakData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/Heroes/RexxarData/RexxarData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/UnitData.xml",
		"mods/heromods/alarak.stormmod/base.stormdata/GameData/AlarakData.xml",
		"mods/heromods/ana.stormmod/base.stormdata/GameData/AnaData.xml",
	} {
		err := x.loadXML(p)
		if err != nil {
			t.Fatal(err)
		}
	}
	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			s, err := x.Get(k)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if s != v {
				t.Fatalf("expected %q, got %q", v, s)
			}
		})
	}
}
