package main

import "testing"

func TestEval(t *testing.T) {
	tests := map[string]float64{
		"-0.75/-0.5-1*100":   50,
		"-0.3--0.15)*(-100)": 15,
		"*-100*-0.5*":        50,
	}
	for tc, expect := range tests {
		t.Run(tc, func(t *testing.T) {
			res := evalExpr(tc)
			if res != expect {
				t.Fatalf("expected %v, got %v", expect, res)
			}
		})
	}
}

func TestXML(t *testing.T) {
	var x XML
	for _, p := range []string{
		"mods/heroesdata.stormmod/base.stormdata/GameData/BehaviorData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/EffectData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/EffectData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/Heroes/AnubarakData/AnubarakData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/Heroes/RexxarData/RexxarData.xml",
		"mods/heroesdata.stormmod/base.stormdata/GameData/UnitData.xml",
		"mods/heromods/alarak.stormmod/base.stormdata/GameData/AlarakData.xml",
		"mods/heromods/ana.stormmod/base.stormdata/GameData/AnaData.xml",
		"mods/heromods/chogall.stormmod/base.stormdata/GameData/ChoGallData.xml",
		"mods/heromods/chromie.stormmod/base.stormdata/GameData/ChromieData.xml",
		"mods/heromods/dva.stormmod/base.stormdata/GameData/DVaData.xml",
		"mods/heromods/stukov.stormmod/base.stormdata/GameData/StukovData.xml",
		"mods/heromods/tracer.stormmod/base.stormdata/GameData/TracerData.xml",
	} {
		err := x.loadXML(p)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Run("extract", func(t *testing.T) {
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
			"Behavior,Artifact_AP_Base,Modification.DamageDealtFraction[Ability]": "0",
			"Behavior,ChromieTimeTrapChronoSicknessSlow,MaxStackCount":            "1",
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
	})

	t.Run("tooltip", func(t *testing.T) {
		tests := map[string]string{
			`Increases Melee's Pulse Bomb charge from <c val="#TooltipNumbers"><d ref="-Effect,TracerMeleeApplyPulseBombChargesHero,Cost[0].ChargeCountUse" player="0"/>%</c> to <c val="#TooltipNumbers"><d ref="-Effect,TracerMeleePulseStrikeApplyPulseBombCharges,Cost[0].ChargeCountUse-Effect,TracerMeleeApplyPulseBombChargesHero,Cost[0].ChargeCountUse" player="0"/>%</c> against Heroes.`: "Increases Melee's Pulse Bomb charge from 10% to 20% against Heroes.",

			`Phase Shifting to an ally casts a free Pixie Dust on them and reveals a large area around them and all enemies in it for <c val="#TooltipNumbers"><d ref="[d ref='Effect,BrightwingPhaseShiftPeekabooCreatePersistentRevealer,PeriodCount' player='0'/] / 16"/></c> seconds.`: "Phase Shifting to an ally casts a free Pixie Dust on them and reveals a large area around them and all enemies in it for 6 seconds.",

			`Locust Basic Attacks cleave for <c val="#TooltipNumbers"><d ref="Effect,AbathurLocustAssaultWeaponCleaveDamage,Amount/Effect,AbathurLocustAssaultWeaponDamage,Amount*100"/>%</c> damage, and explode on death for <c val="#TooltipNumbers"><d ref="Effect,AbathurLocustAssaultExplodeDamage,Amount"/></c> damage.`: "Locust Basic Attacks cleave for 50% damage, and explode on death for 102 damage.",

			`After a <c val="#TooltipNumbers"><d ref="Effect,PrecisionStrikeWarningCreatePersistent,ExpireDelay" precision="1"/></c> second delay, deals <c val="#TooltipNumbers"><d ref="Effect,PrecisionStrikeDamage,Amount*(Behavior,Artifact_AP_Base,Modification.DamageDealtFraction[Ability]+1)"/></c> damage to enemies within an area. Unlimited range.`: "After a 1.5 second delay, deals 456 damage to enemies within an area. Unlimited range.",

			`Shrike can be activated to increase vision radius by <c val="#TooltipNumbers"><d ref="Unit,HeroAna,Sight/Behavior,AnaAimDownSights,Modification.SightBonus*100" player="0" precision="2"/>%</c> and Basic Attack range by <c val="#TooltipNumbers"><d ref="Behavior,AnaAimDownSights,Modification.WeaponRange"  player="0" precision="2"/></c> but reduce Movement Speed by <c val="#TooltipNumbers"><d ref="-Behavior,AnaAimDownSights,Modification.UnifiedMoveSpeedFactor*100" player="0" precision="2"/>%</c>. Lasts until canceled.`: "Shrike can be activated to increase vision radius by 100% and Basic Attack range by 2 but reduce Movement Speed by 20%. Lasts until canceled.",

			`Reduce Time Trap's cooldown and Mana cost by <c val="#TooltipNumbers"><d ref="Talent,ChromieTimeTrapChronoSickness,AbilityModificationArray[0].Modifications[0].Value/Abil,ChromieTimeTrap,Cost.Charge.TimeUse*(-100)" player="0"/>%</c>. After the Time Stop ends, the enemy is also Slowed by <c val="#TooltipNumbers"><d ref="Behavior,ChromieTimeTrapChronoSicknessSlow,MaxStackCount*Behavior,ChromieTimeTrapChronoSicknessSlow,Modification.UnifiedMoveSpeedFactor*(-100)" precision="2"/>%</c> for <c val="#TooltipNumbers"><d ref="Behavior,ChromieTimeTrapChronoSicknessSlow,Duration" precision="2"/></c> seconds.`: "Reduce Time Trap's cooldown and Mana cost by 50%. After the Time Stop ends, the enemy is also Slowed by 50% for 4 seconds.",

			`For <c val="#TooltipNumbers"><d ref="Behavior,DVaPilotNanoweaveSuitBuff,Duration"/></c> seconds after ejecting from her Mech, D.Va gains <c val="#TooltipNumbers"><d ref="Behavior,DVaPilotNanoweaveSuitBuff,ArmorModification.AllArmorBonus"/></c> Armor and her Basic Attacks grant <c val="#TooltipNumbers"><d ref="Effect,DVaPilotNanoweaveSuitWeaponCallMechCooldownReduction,Cost.Cooldown.TimeUse/Effect,DVaPilotWeaponCallMechCooldownReduction,Cost.Cooldown.TimeUse-1*100" precision="2"/>%</c> more cooldown reduction towards Call Mech.`: "For 4 seconds after ejecting from her Mech, D.Va gains 50 Armor and her Basic Attacks grant 50% more cooldown reduction towards Call Mech.",

			`Healing Pathogen can continually spread through Stukov, but its healing is reduced by <c val="#TooltipNumbers"><d ref="*-100*Effect,StukovHealingPathogenHealer,MultiplicativeModifierArray[1].Modifier*" precision="2"/>%</c>.`: "Healing Pathogen can continually spread through Stukov, but its healing is reduced by 50%.",
		}
		for k, v := range tests {
			t.Run(k, func(t *testing.T) {
				s, err := getTooltip(k, x)
				if err != nil {
					t.Fatalf("%+v", err)
				}
				if s != v {
					t.Fatalf("expected %q, got %q", v, s)
				}
			})
		}
	})
}
