package main

type talentText struct {
	Name string
	Text string
}

var talentData = map[string]talentText{
	"AbathurCombatStyleAssaultStrain": {
		Name: "Assault Strain",
		Text: "Locusts cleave and explode",
	},
	"AbathurCombatStyleBombardStrain": {
		Name: "Bombard Strain",
		Text: "Locusts gain a ranged attack",
	},
	"AbathurCombatStyleLocustBrood": {
		Name: "Locust Brood",
		Text: "Spawn a group of Locusts",
	},
	"AbathurCombatStyleSurvivalInstincts": {
		Name: "Survival Instincts",
		Text: "Increases Locust Health and duration",
	},
	"AbathurHeroicAbilityEvolveMonstrosity": {
		Name: "Evolve Monstrosity",
		Text: "Minion or Locust becomes a powerful Monstrosity",
	},
	"AbathurHeroicAbilityUltimateEvolution": {
		Name: "Ultimate Evolution",
		Text: "Clone target allied Hero and control it",
	},
	"AbathurMasteryAdrenalineBoost": {
		Name: "Adrenaline Boost",
		Text: "Carapace increases Move Speed",
	},
	"AbathurMasteryBallistospores": {
		Name: "Ballistospores",
		Text: "Toxic Nests have global range and last longer",
	},
	"AbathurMasteryEnvenomedNestsToxicNest": {
		Name: "Envenomed Nest",
		Text: "Toxic Nests deal more damage",
	},
	"AbathurMasteryEnvenomedSpikes": {
		Name: "Envenomed Spikes",
		Text: "Spike Burst slows Move Speed",
	},
	"AbathurMasteryEvolutionComplete": {
		Name: "Evolution Complete",
		Text: "Monstrosity gains Deep Tunnel",
	},
	"AbathurMasteryEvolutionMasterUltimateEvolution": {
		Name: "Evolution Master",
		Text: "Evolution Master Lower the cooldown of Ultimate Evolution to 90 seconds, and increase the duration to 60 seconds.",
	},
	"AbathurMasteryLocustMaster": {
		Name: "Locust Nest",
		Text: "Create Locust Nests",
	},
	"AbathurMasteryNeedlespine": {
		Name: "Needlespine",
		Text: "Increases Stab damage and range",
	},
	"AbathurMasteryPressurizedGlands": {
		Name: "Pressurized Glands",
		Text: "Increases Spike Burst range and decreases cooldown",
	},
	"AbathurMasteryProlificDispersal": {
		Name: "Prolific Dispersal",
		Text: "Increases Toxic Nest charges",
	},
	"AbathurMasteryRegenerativeMicrobes": {
		Name: "Regenerative Microbes",
		Text: "Carapace heals target",
	},
	"AbathurMasterySpatialEfficiency": {
		Name: "Spatial Efficiency",
		Text: "Stab gains an additional charge",
	},
	"AbathurMasteryVileNestsToxicNest": {
		Name: "Vile Nest",
		Text: "Toxic Nests slow Move Speed",
	},
	"AbathurSymbioteAdrenalOverload": {
		Name: "Adrenal Overload",
		Text: "Symbiote host gains Attack Speed",
	},
	"AbathurSymbioteCarapaceNetworkedCarapace": {
		Name: "Networked Carapace",
		Text: "Carapace shields all nearby allies",
	},
	"AbathurSymbioteCarapaceSustainedCarapace": {
		Name: "Sustained Carapace",
		Text: "Increases Carapace shield and persists after Symbiote",
	},
	"AbathurSymbioteHivemind": {
		Name: "Hivemind",
		Text: "Symbiote two targets",
	},
	"AbathurSymbioteSpikeBurstSomaTransference": {
		Name: "Soma Transference",
		Text: "Spike Burst heals for each Hero hit",
	},
	"AbathurUltimateEvolutionEvolutionaryLink": {
		Name: "Evolutionary Link",
		Text: "Ultimate Evolution target gets a Shield",
	},
	"AbathurVolatileMutation": {
		Name: "Volatile Mutation",
		Text: "Heroic summons radiate damage",
	},
	"AlarakAppliedForce": {
		Name: "Applied Force",
		Text: "Reduce Sadism, increase Telekinesis distance, range",
	},
	"AlarakBladeoftheHighlord": {
		Name: "Blade of the Highlord",
		Text: "Basic Attacks increase Sadism",
	},
	"AlarakChaosReigns": {
		Name: "Chaos Reigns",
		Text: "Quest: Hit Heroes to increase Discord Strike damage",
	},
	"AlarakCounterStrikeItem": {
		Name: "Counter-Strike",
		Text: "Prevents damage to deal damage in a large area",
	},
	"AlarakDeadlyChargeItem": {
		Name: "Deadly Charge",
		Text: "Channel to charge a long distance",
	},
	"AlarakDissonance": {
		Name: "Dissonance",
		Text: "Increase Discord Strike's Silence duration",
	},
	"AlarakExtendedLightning": {
		Name: "Extended Lightning",
		Text: "Quest: Reduce Sadism, empower Lightning Surge",
	},
	"AlarakHastyBargain": {
		Name: "Hasty Bargain",
		Text: "Activate to reduce Sadism bonus and reset Basic Ability cooldowns",
	},
	"AlarakHeroicAbilityCounterStrike": {
		Name: "Counter-Strike",
		Text: "Prevents damage to deal damage in a large area",
	},
	"AlarakHeroicAbilityDeadlyCharge": {
		Name: "Deadly Charge",
		Text: "Channel to charge a long distance",
	},
	"AlarakHinderedMotion": {
		Name: "Hindered Motion",
		Text: "Telekinesis Slows",
	},
	"AlarakLastLaugh": {
		Name: "Last Laugh",
		Text: "Cleanse and teleport, but reduce Health",
	},
	"AlarakLethalOnslaughtDiscordStrike": {
		Name: "Lethal Onslaught",
		Text: "Discord Strike grants Sadism bonus to Basic Attacks",
	},
	"AlarakLightningBarrage": {
		Name: "Lightning Barrage",
		Text: "Center hits grant free Lightning Surge cast",
	},
	"AlarakMockingStrikes": {
		Name: "Mocking Strikes",
		Text: "Reduce cooldowns against disabled Heroes",
	},
	"AlarakNegativelyCharged": {
		Name: "Negatively Charged",
		Text: "Quest: Increase Lightning Surge center damage",
	},
	"AlarakPureMalice": {
		Name: "Pure Malice",
		Text: "Allied deaths increase Sadism",
	},
	"AlarakRiteofRakShir": {
		Name: "Rite of Rak'Shir",
		Text: "Mark and damage Hero to increase Sadism",
	},
	"AlarakRuthlessMomentum": {
		Name: "Ruthless Momentum",
		Text: "Reduce cooldowns at high Health",
	},
	"AlarakShowofForce": {
		Name: "Show of Force",
		Text: "Ability combos deal bonus damage",
	},
	"AlarakSustainingPower": {
		Name: "Sustaining Power",
		Text: "Increase Lightning Surge healing",
	},
	"AmazonBallLightning": {
		Name: "Ball Lightning",
		Text: "Throw a bouncing lightning ball",
	},
	"AmazonChargedStrikes": {
		Name: "Charged Strikes",
		Text: "Activate to empower Basic Attacks",
	},
	"AmazonGroundingBolt": {
		Name: "Grounding Bolt",
		Text: "Lightning Fury Slows",
	},
	"AmazonImpale": {
		Name: "Impale",
		Text: "Increase Fend damage",
	},
	"AmazonImprisoningLight": {
		Name: "Imprisoning Light",
		Text: "Valkyrie Roots enemies",
	},
	"AmazonInfiniteLightning": {
		Name: "Infinite Lightning",
		Text: "Ball Lightning bounces indefinitely",
	},
	"AmazonInnerLight": {
		Name: "Inner Light",
		Text: "Emit Blinding Light when Stunned or Rooted",
	},
	"AmazonLungingStrike": {
		Name: "Lunging Strike",
		Text: "Increase Fend area, duration",
	},
	"AmazonMartialLaw": {
		Name: "Martial Law",
		Text: "Disabled Heroes take additional damage",
	},
	"AmazonPenetrate": {
		Name: "Penetrate",
		Text: "Consecutive Fend hits deal increased damage",
	},
	"AmazonPierce": {
		Name: "Pierce",
		Text: "Lightning Fury pierces",
	},
	"AmazonPlateoftheWhale": {
		Name: "Plate of the Whale",
		Text: "Quest: Avoidance regenerates Health",
	},
	"AmazonRingoftheLeech": {
		Name: "Ring of the Leech",
		Text: "Damaging Blinded targets heals",
	},
	"AmazonSeraphsHymn": {
		Name: "Seraph's Hymn",
		Text: "Basic Attacks reduce Blinding Light cooldown",
	},
	"AmazonSurgeOfLight": {
		Name: "Surge of Light",
		Text: "Activate Avoidance to deal damage",
	},
	"AmazonThundergodsVigor": {
		Name: "Thundergod's Vigor",
		Text: "Reduce Lightning Fury cooldown",
	},
	"AmazonThunderstroke": {
		Name: "Thunderstroke",
		Text: "Quest: Empower Lightning Fury",
	},
	"AmazonTitansRevenge": {
		Name: "Titan's Revenge",
		Text: "Basic Attacks ignore Armor",
	},
	"AmazonTrueSight": {
		Name: "True Sight",
		Text: "Quest: Empower Blinding Light",
	},
	"AmazonValkyrie": {
		Name: "Valkyrie",
		Text: "Summon a Valkyrie to impale enemies",
	},
	"AmazonWarTraveler": {
		Name: "War Traveler",
		Text: "Avoidance grants Movement Speed",
	},
	"AnaAimDownSights": {
		Name: "Aim Down Sights",
		Text: "Activate Shrike to increase vision, range",
	},
	"AnaAimDownSightsCustomOptics": {
		Name: "Custom Optics",
		Text: "Increase Aim Down Sights range bonus",
	},
	"AnaBioticGrenadeAirStrike": {
		Name: "Air Strike",
		Text: "Activate to throw Biotic Grenade farther",
	},
	"AnaBioticGrenadeContactHealing": {
		Name: "Contact Healing",
		Text: "Increase Biotic Grenade healing on multiple Heroes",
	},
	"AnaBioticGrenadeGrenadeCalibration": {
		Name: "Grenade Calibration",
		Text: "Quest: Empower Biotic Grenade",
	},
	"AnaDebilitatingDart": {
		Name: "Debilitating Dart",
		Text: "Fire a damage-reducing dart",
	},
	"AnaDetachableBoxMagazine": {
		Name: "Detachable Box Magazine",
		Text: "Quest: Empower Doses",
	},
	"AnaDynamicShooting": {
		Name: "Dynamic Shooting",
		Text: "Basic Attacks increase Attack Speed",
	},
	"AnaEyeOfHorusBallisticAdvantage": {
		Name: "Ballistic Advantage",
		Text: "Eye of Horus rounds explode",
	},
	"AnaHealingDartConcentratedDoses": {
		Name: "Concentrated Doses",
		Text: "Doses empower Healing Dart",
	},
	"AnaHealingDartPurifyingDarts": {
		Name: "Purifying Darts",
		Text: "Healing Dart removes Roots, Slows",
	},
	"AnaHealingDartSharpshooter": {
		Name: "Sharpshooter",
		Text: "Empower consecutive Healing Darts",
	},
	"AnaHealingDartSmellingSalts": {
		Name: "Smelling Salts",
		Text: "Healing Dart removes Stuns",
	},
	"AnaHealingDartSpeedSerum": {
		Name: "Speed Serum",
		Text: "Healing Dart grants Move Speed",
	},
	"AnaHeroicAbilityEyeOfHorus": {
		Name: "Eye of Horus",
		Text: "Fire unlimited range healing and damaging rounds",
	},
	"AnaHeroicAbilityNanaBoost": {
		Name: "Nano Boost",
		Text: "Grant ally Spell Power, reduced cooldowns",
	},
	"AnaMindNumbingAgent": {
		Name: "Mind-Numbing Agent",
		Text: "Doses decrease Spell Power",
	},
	"AnaNanoInfusion": {
		Name: "Nano Infusion",
		Text: "Nano Boost heals for damage dealt",
	},
	"AnaOverdose": {
		Name: "Overdose",
		Text: "Sleep Dart applies Doses",
	},
	"AnaPiercingDarts": {
		Name: "Piercing Darts",
		Text: "Quest: Empower Sleep Dart, Healing Dart",
	},
	"AnaSleepingDartTemporaryBlindness": {
		Name: "Temporary Blindness",
		Text: "Sleep Dart Blinds",
	},
	"AnaSomnolentDoses": {
		Name: "Somnolent Doses",
		Text: "Max Dosage puts enemies to Sleep",
	},
	"AnubarakAcidDrenchedMandibles": {
		Name: "Acid Drenched Mandibles",
		Text: "Attacking disabled Heroes increases damage",
	},
	"AnubarakCocoonCryptweave": {
		Name: "Cryptweave",
		Text: "Increases Cocoon duration",
	},
	"AnubarakCombatStyleChitinousPlating": {
		Name: "Chitinous Plating",
		Text: "Reduces Harden Carapace cooldown",
	},
	"AnubarakCombatStyleLegionOfBeetles": {
		Name: "Legion of Beetles",
		Text: "Spawn Beetles periodically",
	},
	"AnubarakDebilitation": {
		Name: "Debilitation",
		Text: "Burrow Charge decreases enemy Spell Power",
	},
	"AnubarakHeroicAbilityCarrionSwarm": {
		Name: "Locust Swarm",
		Text: "Damages enemies while healing",
	},
	"AnubarakHeroicAbilityCocoon": {
		Name: "Cocoon",
		Text: "Wraps enemy in a cocoon",
	},
	"AnubarakMasteryBedOfBarbs": {
		Name: "Bed of Barbs",
		Text: "Impale slows and damages over time",
	},
	"AnubarakMasteryEpicenterBurrowCharge": {
		Name: "Epicenter",
		Text: "Increases Burrow Charge area and reduces cooldown",
	},
	"AnubarakMasteryHiveMasterCarrionSwarm": {
		Name: "Hive Master",
		Text: "Gain a permanent life-stealing Locust",
	},
	"AnubarakMasteryLeechingScarabs": {
		Name: "Leeching Scarabs",
		Text: "Beetles heal Anub'arak",
	},
	"AnubarakMasteryPersistentCarapaceHardenCarapace": {
		Name: "Persistent Carapace",
		Text: "Longer lasting Shield",
	},
	"AnubarakMasteryShedExoskeletonHardenCarapace": {
		Name: "Shed Exoskeleton",
		Text: "Harden Carapace increases Move Speed",
	},
	"AnubarakMasteryUnderkingBurrowCharge": {
		Name: "Underking",
		Text: "Increases Burrow Charge range and damage",
	},
	"AnubarakMasteryUrticatingSpines": {
		Name: "Urticating Spines",
		Text: "Harden Carapace deals area damage",
	},
	"AnubarakNerubianArmor": {
		Name: "Nerubian Armor",
		Text: "Periodically gain Spell Armor",
	},
	"AnubarakResilientScarabs": {
		Name: "Resilient Scarabs",
		Text: "Beetles gain Spell Armor",
	},
	"AnubarakScarabHostBeetleJuiced": {
		Name: "Beetle, Juiced",
		Text: "Basic Attacks to Heroes spawn Beetles",
	},
	"AnubarakSubterraneanShield": {
		Name: "Subterranean Shield",
		Text: "Burrow Charge grants Shield",
	},
	"ArtanisBladeDashPsionicSynergy": {
		Name: "Psionic Synergy",
		Text: "Increases Blade Dash cooldown reduction of Shield Overload",
	},
	"ArtanisBladeDashSolariteReaper": {
		Name: "Solarite Reaper",
		Text: "Increases Blade Dash damage",
	},
	"ArtanisBladeDashTemplarsZeal": {
		Name: "Templar's Zeal",
		Text: "Blade Dash recharges faster when at low Health",
	},
	"ArtanisHeroicAbilitySpearofAdunPurifierBeam": {
		Name: "Purifier Beam",
		Text: "Chase an enemy with an orbital beam",
	},
	"ArtanisHeroicAbilitySpearofAdunSuppressionPulse": {
		Name: "Suppression Pulse",
		Text: "Damage and Blind enemies in a large area",
	},
	"ArtanisPhasePrismChronoSurge": {
		Name: "Chrono Surge",
		Text: "Phase Prism grants Attack Speed",
	},
	"ArtanisPhasePrismGravitonVortex": {
		Name: "Graviton Vortex",
		Text: "Reduces Phase Prism cooldown and pull another Hero",
	},
	"ArtanisPhasePrismWarpSickness": {
		Name: "Warp Sickness",
		Text: "Phase Prism slows enemies",
	},
	"ArtanisPlasmaBurn": {
		Name: "Plasma Burn",
		Text: "Shield Overload burns nearby enemies",
	},
	"ArtanisShieldOverloadForceofWill": {
		Name: "Force of Will",
		Text: "Increases Shield Overload Basic Attack reduction",
	},
	"ArtanisShieldOverloadPhaseBulwark": {
		Name: "Phase Bulwark",
		Text: "Shield Overload grants Spell Armor",
	},
	"ArtanisShieldOverloadShieldBattery": {
		Name: "Shield Battery",
		Text: "Shield Overload recharges faster while active",
	},
	"ArtanisShieldOverloadShieldSurge": {
		Name: "Shield Surge",
		Text: "Increase Shield Overload shield at low Health",
	},
	"ArtanisSpearofAdunPurifierBeamTargetPurified": {
		Name: "Target Purified",
		Text: "Purifier Beam recasts if target dies",
	},
	"ArtanisSpearofAdunSuppressionPulseOrbitalBombardment": {
		Name: "Orbital Bombardment",
		Text: "Suppression Pulse gains an additional charge",
	},
	"ArtanisTwinBladesAmateurOpponent": {
		Name: "Amateur Opponent",
		Text: "Increases Twin Blades damage to non-Heroes",
	},
	"ArtanisTwinBladesPsionicWound": {
		Name: "Psionic Wound",
		Text: "Twin Blades lowers Armor",
	},
	"ArtanisTwinBladesReactiveParry": {
		Name: "Reactive Parry",
		Text: "Twin Blades gives Physical Armor",
	},
	"ArtanisTwinBladesTitanKiller": {
		Name: "Titan Killer",
		Text: "Twin Blades deals more damage versus Heroes",
	},
	"ArtanisTwinBladesTripleStrike": {
		Name: "Triple Strike",
		Text: "Twin Blades strikes 3 times but increases cooldown",
	},
	"ArtanisTwinBladesZealotCharge": {
		Name: "Zealot Charge",
		Text: "Increase Twin Blades charge distance",
	},
	"ArthasAntiMagicShell": {
		Name: "Anti-Magic Shell",
		Text: "Activate to make Arthas immune to Spell Damage for  seconds and heal Arthas for 25% of the damage prevented.",
	},
	"ArthasDeathlord": {
		Name: "Deathlord",
		Text: "Reduces Death Coil's cooldown by  seconds and increase its range by 30%.",
	},
	"ArthasDeathsAdvance": {
		Name: "Death's Advance",
		Text: "Activate to increase Movement Speed by % for 3 seconds.Passive: Increases Movement Speed by %.",
	},
	"ArthasHeroicAbilityArmyoftheDead": {
		Name: "Army of the Dead",
		Text: "Raise ghouls that fight enemies, heal Arthas",
	},
	"ArthasHeroicAbilitySummonSindragosa": {
		Name: "Summon Sindragosa",
		Text: "Disables enemies and Structures",
	},
	"ArthasIceboundFortitude": {
		Name: "Icebound Fortitude",
		Text: "Activate to gain  Armor, reducing damage taken by %, and reduce the duration of Stuns, Slows, and Roots against Arthas by 75% for  seconds.",
	},
	"ArthasIcyTalons": {
		Name: "Icy Talons",
		Text: "Frozen Tempest increases Attack Speed",
	},
	"ArthasMasteryAbsoluteZeroSummonSindragosa": {
		Name: "Absolute Zero",
		Text: "Increases Sindragosa range and disable duration",
	},
	"ArthasMasteryBitingColdFrozenTempest": {
		Name: "Biting Cold",
		Text: "Frozen Tempest damage increased",
	},
	"ArthasMasteryEmbraceDeathDeathCoil": {
		Name: "Embrace Death",
		Text: "Death Coil deals more damage when missing Health",
	},
	"ArthasMasteryEternalHungerFrostmourneHungers": {
		Name: "Eternal Hunger",
		Text: "Quest: Empower Frostmourne Hungers",
	},
	"ArthasMasteryFrostPresenceHowlingBlast": {
		Name: "Frost Presence",
		Text: "Quest: Root Heroes to empower Howling Blast",
	},
	"ArthasMasteryFrostStrikeFrostmourneHungers": {
		Name: "Frost Strike",
		Text: "Reduces Frostmourne Hungers cooldown and adds slow",
	},
	"ArthasMasteryFrostmourneFeedsFrostmourneHungers": {
		Name: "Frostmourne Feeds",
		Text: "Increases number of Frostmourne Hungers attacks",
	},
	"ArthasMasteryFrozenWastesFrozenTempest": {
		Name: "Frozen Wastes",
		Text: "Quest: Frozen Tempest slow lingers",
	},
	"ArthasMasteryImmortalCoilDeathCoil": {
		Name: "Immortal Coil",
		Text: "Death Coil always heals",
	},
	"ArthasMasteryLegionOfNorthrendArmyoftheDead": {
		Name: "Legion of Northrend",
		Text: "Ghouls last longer and more are created",
	},
	"ArthasMasteryRemorselessWinterFrozenTempest": {
		Name: "Remorseless Winter",
		Text: "Enemy Heroes that remain within Frozen Tempest for 3 seconds are rooted for  seconds. This effect can only happen once every  seconds.",
	},
	"ArthasRime": {
		Name: "Rime",
		Text: "Periodically gain Physical Armor",
	},
	"ArthasRuneTap": {
		Name: "Rune Tap",
		Text: "Consecutive Basic Attacks heal",
	},
	"ArthasShatteredArmor": {
		Name: "Shattered Armor",
		Text: "Howling Blast reduces Armor",
	},
	"ArtifactAbilityPower": {
		Name: "Brilliant Topaz",
		Text: "Increases ability damage and healing",
	},
	"ArtifactAttackDamage": {
		Name: "Primal Ruby",
		Text: "Increases basic attack damage",
	},
	"ArtifactAttackSpeed": {
		Name: "Skyfire Emerald",
		Text: "Increases speed of basic attacks",
	},
	"ArtifactCooldownReduction": {
		Name: "Royal Diamond",
		Text: "Reduces cooldowns for abilities",
	},
	"ArtifactHealthRegen": {
		Name: "Warrior's Resolve",
		Text: "Increases Health regeneration",
	},
	"ArtifactLifesteal": {
		Name: "Blood Siphon",
		Text: "Basic attacks restore Health",
	},
	"ArtifactManaRegen": {
		Name: "Mana Infusion",
		Text: "Increases Mana regeneration",
	},
	"ArtifactMaxHealth": {
		Name: "Bold Amethyst",
		Text: "Increases maximum Health",
	},
	"ArtifactMaxMana": {
		Name: "Khaydarin Amulet",
		Text: "Increases maximum Mana",
	},
	"ArtifactMovementSpeed": {
		Name: "Wirt's Leg",
		Text: "Increases movement speed",
	},
	"ArtifactShields": {
		Name: "Nexus Sapphire",
		Text: "Grants shields",
	},
	"ArtifactSiegeDamage": {
		Name: "Destroyer",
		Text: "Increases damage vs. minions and towers",
	},
	"AurielAngelicFlightTalent": {
		Name: "Angelic Flight",
		Text: "After  seconds, fly to a target location.",
	},
	"AurielBlindingFlashSacredSweepTalent": {
		Name: "Blinding Flash",
		Text: "Enemies hit by the center area of Sacred Sweep are blinded for  seconds.",
	},
	"AurielBurstingLightRayOfHeavenTalent": {
		Name: "Bursting Light",
		Text: "Reduces the cooldown of Ray of Heaven by  seconds.",
	},
	"AurielConvergingForceSacredSweepTalent": {
		Name: "Converging Force",
		Text: "Enemies hit by the outer area are pushed slightly toward the center.",
	},
	"AurielDiamondResolveCrystalAegisTalent": {
		Name: "Diamond Resolve",
		Text: "When Crystal Aegis expires, it grants the target  Armor for  seconds, reducing damage taken by %.",
	},
	"AurielEmpathicLinkBestowHopeTalent": {
		Name: "Empathic Link",
		Text: "Auriel stores 25% of damage taken by allies with Bestow Hope.",
	},
	"AurielEnergizedCordRayOfHeavenTalent": {
		Name: "Energized Cord",
		Text: "Increases the energy stored from Auriel's Basic Attacks to % of the damage against Heroes and % of the damage against non-Heroes.Does not affect Auriel's Bestow Hope ally.",
	},
	"AurielGlimmerofHopeRayOfHeavenTalent": {
		Name: "Glimmer of Hope",
		Text: "Collecting a Regeneration Globe reduces the cost of Auriel's next Ray of Heaven by 75%.",
	},
	"AurielHeavyBurdenDetainmentStrikeTalent": {
		Name: "Heavy Burden",
		Text: "When an enemy Hero is stunned by Detainment Strike, their Movement Speed is slowed by % for  seconds after the stun.",
	},
	"AurielHeroicCrystalAegis": {
		Name: "Crystal Aegis",
		Text: "Put an allied Hero into a Stasis crystal that explodes",
	},
	"AurielHeroicResurrect": {
		Name: "Resurrect",
		Text: "Channel on the spirit of a dead ally for  seconds. After a 5 second delay, they are brought back to life with 50% of their maximum Health at the location where they died.",
	},
	"AurielIncreasingClaritySacredSweepTalent": {
		Name: "Increasing Clarity",
		Text: "Quest: Every time Sacred Sweep hits a Hero in the center, increase the center damage by , up to . Reward: After hitting  Heroes, this center damage bonus is increased to .",
	},
	"AurielLightSpeedResurrectTalent": {
		Name: "Light Speed",
		Text: "Resurrected allies gain % increased Movement Speed, decaying over  seconds.  While a resurrected ally remains alive, Resurrect's next cooldown recharges % faster.",
	},
	"AurielMajesticSpanSacredSweepTalent": {
		Name: "Majestic Span",
		Text: "Increases the radius of Sacred Sweep by %.",
	},
	"AurielPiercingLashDetainmentStrikeTalent": {
		Name: "Piercing Lash",
		Text: "Detainment Strike now pierces and hits all enemy Heroes in a line, reducing the cooldown by  seconds for each Hero hit.",
	},
	"AurielRepeatedOffenseDetainmentStrikeTalent": {
		Name: "Repeated Offense",
		Text: "Quest: Every time Detainment Strike stuns a Hero, increase the stun damage by , up to .Reward: After stunning  Heroes, increase this damage bonus to .",
	},
	"AurielRepellingStrikeDetainmentStrikeTalent": {
		Name: "Repelling Strike",
		Text: "Enemies hit by Detainment Strike are knocked back 35% farther.",
	},
	"AurielReservoirofHopeRayOfHeavenTalent": {
		Name: "Reservoir of Hope",
		Text: "Quest: Each maximum energy Ray of Heaven Auriel casts increases the maximum amount of energy that can be stored by .",
	},
	"AurielRighteousAssaultSacredSweepTalent": {
		Name: "Righteous Assault",
		Text: "Reduces the cooldown of Sacred Sweep by  seconds for each enemy Hero hit by its center.",
	},
	"AurielSearingLightRayOfHeavenTalent": {
		Name: "Searing Light",
		Text: "Ray of Heaven also deals damage to enemies in the area equal to 30% of the energy consumed.",
	},
	"AurielShieldOfHopeTalent": {
		Name: "Shield of Hope",
		Text: "Activate to grant all nearby allied Heroes a shield for 3 seconds equal to 50% of the amount of Health they are missing.",
	},
	"AurielSwiftSweepSacredSweepTalent": {
		Name: "Swift Sweep",
		Text: "Increases the cast speed of Sacred Sweep by 50%.",
	},
	"AurielWillofHeavenBestowHopeTalent": {
		Name: "Will of Heaven",
		Text: "Allies with Bestow Hope gain % Attack Speed.",
	},
	"AurielWrathofHeavenBestowHopeTalent": {
		Name: "Wrath of Heaven",
		Text: "Allies with Bestow Hope gain % Spell Power.",
	},
	"AzmodanBloodForBlood": {
		Name: "Blood for Blood",
		Text: "Activate to steal Health from your target",
	},
	"AzmodanBoltoftheStorms": {
		Name: "Bolt of the Storm",
		Text: "Activate to teleport a short distance",
	},
	"AzmodanBombardment": {
		Name: "Bombardment",
		Text: "[PH] Reduce cooldown on Hero kill",
	},
	"AzmodanBribe": {
		Name: "Bribe",
		Text: "Kill Minions to bribe a Mercenary",
	},
	"AzmodanGluttonousWard": {
		Name: "Gluttonous Ward",
		Text: "Place a ward that restores Azmodan's Health and Mana",
	},
	"AzmodanHellishHirelings": {
		Name: "Hellish Hirelings",
		Text: "General of Hell improves Mercenary damage",
	},
	"AzmodanHeroicAbilityBlackPool": {
		Name: "Black Pool",
		Text: "Create an empowering zone",
	},
	"AzmodanHeroicAbilityDemonicInvasion": {
		Name: "Demonic Invasion",
		Text: "Damage area and create an army of Grunts",
	},
	"AzmodanMasteryArmyOfHell": {
		Name: "Army of Hell",
		Text: "Demon Warriors deal more damage, cost less Mana",
	},
	"AzmodanMasteryBattleborn": {
		Name: "Battleborn",
		Text: "Globe of Annihilation summons a Demon Warrior",
	},
	"AzmodanMasteryBlackPoolFifthCircle": {
		Name: "Fifth Circle",
		Text: "Black Pool gives slowing attacks",
	},
	"AzmodanMasteryBoundMinion": {
		Name: "Bound Minion",
		Text: "General of Hell improves a Minion",
	},
	"AzmodanMasteryDemonicSmite": {
		Name: "Demonic Smite",
		Text: "Demon Lieutenants cast Demonic Smite",
	},
	"AzmodanMasteryEnduringWarriors": {
		Name: "Enduring Warriors",
		Text: "Demons last 5 seconds longer and do an additional % attack damage.",
	},
	"AzmodanMasteryForcedRecruitment": {
		Name: "Forced Recruitment",
		Text: "Reduces General of Hell cooldown, gains a charge",
	},
	"AzmodanMasteryGluttony": {
		Name: "Gluttony",
		Text: "All Shall Burn heals Azmodan",
	},
	"AzmodanMasteryHedonism": {
		Name: "Hedonism",
		Text: "Reduces Globe of Annihilation Mana cost",
	},
	"AzmodanMasteryHellforgedArmor": {
		Name: "Hellforged Armor",
		Text: "Demon Warriors burn enemies and take less damage",
	},
	"AzmodanMasteryInfernalGlobe": {
		Name: "Infernal Globe",
		Text: "Reduces cast time and increases missile speed",
	},
	"AzmodanMasteryInfusedPower": {
		Name: "Infused Power",
		Text: "Increases All Shall Burn's damage",
	},
	"AzmodanMasteryMarchOfSin": {
		Name: "March of Sin",
		Text: "Increases All Shall Burn Movement Speed",
	},
	"AzmodanMasteryMasterOfDestruction": {
		Name: "Master of Destruction",
		Text: "Reduces All Shall Burn Mana cost, increases Range",
	},
	"AzmodanMasteryPerishingFlame": {
		Name: "Perishing Flame",
		Text: "Demonic Grunts explode on death",
	},
	"AzmodanMasterySiegingWrath": {
		Name: "Sieging Wrath",
		Text: "Quest: Hitting Heroes empowers Globe of Annihilation",
	},
	"AzmodanMasteryTasteForBlood": {
		Name: "Taste for Blood",
		Text: "Quest: Kills increase Globe of Annihilation damage",
	},
	"AzmodanSinforSin": {
		Name: "Sin for Sin",
		Text: "Activate to steal Health from an enemy",
	},
	"AzmodanSinsGrasp": {
		Name: "Sin's Grasp",
		Text: "Activate to deal damage over time",
	},
	"BarbarianAmplifiedHealing": {
		Name: "Amplified Healing",
		Text: "All healing effects increased",
	},
	"BarbarianCombatStyleEndlessFury": {
		Name: "Endless Fury",
		Text: "Increases max Fury",
	},
	"BarbarianCombatStyleFerociousHealing": {
		Name: "Ferocious Healing",
		Text: "Spend Fury to gain Health",
	},
	"BarbarianCombatStyleNoEscape": {
		Name: "No Escape",
		Text: "Abilities give more Movement Speed",
	},
	"BarbarianCombatStyleShotofFury": {
		Name: "Shot of Fury",
		Text: "Increases maximum Fury and can instantly gain more Fury",
	},
	"BarbarianHeroicAbilityLeap": {
		Name: "Leap",
		Text: "Jump to a location, deal damage, and stun enemies",
	},
	"BarbarianHeroicAbilityWrathoftheBerserker": {
		Name: "Wrath of the Berserker",
		Text: "Increase damage and reduce disable duration",
	},
	"BarbarianMasteryAftershock": {
		Name: "Aftershock",
		Text: "Using Seismic Slam temporarily reduces its Fury cost",
	},
	"BarbarianMasteryAngerManagementWrathoftheBerserker": {
		Name: "Anger Management",
		Text: "Increases Wrath Fury generation",
	},
	"BarbarianMasteryArreatCraterLeap": {
		Name: "Arreat Crater",
		Text: "Leap creates a crater and has reduced cooldown",
	},
	"BarbarianMasteryBoonOfTheAncients": {
		Name: "Boon Of The Ancients",
		Text: "Boon Of The Ancients Hitting an enemy reduces Ancient Spear's cooldown by 5 seconds.",
	},
	"BarbarianMasteryCompositeSpearAncientSpear": {
		Name: "Composite Spear",
		Text: "Increases Ancient Spear range and Fury generation",
	},
	"BarbarianMasteryCycloneWhirlwind": {
		Name: "Cyclone",
		Text: "Increases the radius of Whirlwind by 80%.",
	},
	"BarbarianMasteryDustDevils": {
		Name: "Dust Devils",
		Text: "Whirlwind creates Tornadoes that do an additional 25% damage.",
	},
	"BarbarianMasteryEnduringWhirlwind": {
		Name: "Enduring Whirlwind",
		Text: "Whirlwind duration increased by 50%.",
	},
	"BarbarianMasteryFuriousBlowSeismicSlam": {
		Name: "Furious Blow",
		Text: "Increases Seismic Slam damage and cost",
	},
	"BarbarianMasteryHamstringSpiralWhirlwind": {
		Name: "Hamstring Spiral",
		Text: "Whirlwind slows enemy Movement Speed by % for  second.",
	},
	"BarbarianMasteryHurricaneWhirlwind": {
		Name: "Hurricane",
		Text: "Whirlwind removes slows and roots",
	},
	"BarbarianMasteryLifeFunnelWhirlwind": {
		Name: "Life Funnel",
		Text: "Increases Whirlwind healing",
	},
	"BarbarianMasteryMysticalSpearAncientSpear": {
		Name: "Mystical Spear",
		Text: "Reduces Ancient Spear cooldown and always pulls",
	},
	"BarbarianMasteryPoisonedSpearAncientSpear": {
		Name: "Poisoned Spear",
		Text: "Increases Ancient Spear damage",
	},
	"BarbarianMasteryShatteredGroundSeismicSlam": {
		Name: "Shattered Ground",
		Text: "Increases Seismic Slam splash",
	},
	"BarbarianMasteryStrengthFromEarthSeismicSlam": {
		Name: "Strength from Earth",
		Text: "Strength From Earth Seismic Slam heals for % of damage dealt.",
	},
	"BarbarianMasteryWarPaint": {
		Name: "War Paint",
		Text: "Basic Attacks heal",
	},
	"BattleMomentumCrusader": {
		Name: "Blessed Momentum",
		Text: "Basic Attacks reduce Ability cooldowns",
	},
	"BattleMomentumETC": {
		Name: "Battle Momentum",
		Text: "Basic Attacks decrease Ability cooldowns",
	},
	"BattleMomentumKerrigan": {
		Name: "Bladed Momentum",
		Text: "Basic Attacks reduce Ability cooldowns",
	},
	"BattleMomentumMuradin": {
		Name: "Iron-forged Momentum",
		Text: "Basic Attacks reduce Ability cooldowns",
	},
	"BattleMomentumNova": {
		Name: "Battle Momentum",
		Text: "Basic Attacks decrease Ability cooldowns",
	},
	"BattleMomentumRehgar": {
		Name: "Battle Momentum",
		Text: "Basic Attacks decrease Ability cooldowns",
	},
	"BattleMomentumSylvanas": {
		Name: "Battle Momentum",
		Text: "Basic Attacks decrease Ability cooldowns",
	},
	"BattleMomentumTyrael": {
		Name: "Angelic Momentum",
		Text: "Basic Attacks reduce Ability cooldowns",
	},
	"BrightwingArcaneBarrageArcaneFlare": {
		Name: "Arcane Barrage",
		Text: "Increase Arcane Flare range",
	},
	"BrightwingBouncyDustPixieDust": {
		Name: "Bouncy Dust",
		Text: "Pixie Dust bounces",
	},
	"BrightwingDoubleWyrmholeBlinkHeal": {
		Name: "Double Wyrmhole",
		Text: "Free Blink Heal on second target",
	},
	"BrightwingDreamShotArcaneFlare": {
		Name: "Dream Shot",
		Text: "Empowers Arcane Flare",
	},
	"BrightwingGreaterPolymorphPolymorph": {
		Name: "Greater Polymorph",
		Text: "Increases Polymorph duration",
	},
	"BrightwingHyperShiftPhaseShift": {
		Name: "Hyper Shift",
		Text: "Soothing Mist reduces Phase Shift cooldown",
	},
	"BrightwingManicPixiePixieDust": {
		Name: "Manic Pixie",
		Text: "Pixie Dust increases Soothing Mist healing",
	},
	"BrightwingMistifiedSoothingMist": {
		Name: "Mistified",
		Text: "Basic Abilities reduce Soothing Mist cooldown",
	},
	"BrightwingPeekabooPhaseShift": {
		Name: "Peekaboo!",
		Text: "Empowers Phase Shift",
	},
	"BrightwingPixieBoostPixieDust": {
		Name: "Pixie Boost",
		Text: "Pixie Dust grants an additional burst of speed",
	},
	"BrightwingPixieCharm": {
		Name: "Pixie Charm",
		Text: "Heal Allies with Soothing Mist to bribe a Mercenary",
	},
	"BrightwingRevitalizingMistSoothingMist": {
		Name: "Revitalizing Mist",
		Text: "Soothing Mist heals more subsequently",
	},
	"BrightwingUnfurlingSpraySoothingMist": {
		Name: "Unfurling Spray",
		Text: "Increases Soothing Mist range",
	},
	"BrightwingUnstableAnomalyPolymorph": {
		Name: "Unstable Anomaly",
		Text: "Polymorphed targets explode",
	},
	"BrightwingWoundedAnimalPolymorph": {
		Name: "Wounded Animal",
		Text: "Increase the Movement Speed slow of Polymorph to 30%.",
	},
	"BucketCallOfTheStorm": {
		Name: "Onslaught of the Storm",
		Text: "Onslaught of the StormEnemies killed by Abilities explode for  damage in a small area.",
	},
	"BucketConstantHeroism": {
		Name: "Constant Heroism",
		Text: "Constant Heroism Your Heroic Ability's cooldown is reduced by 30%.",
	},
	"BucketFuryOfTheStorm": {
		Name: "Fury of the Storm",
		Text: "Basic Attacks deal damage to nearby enemies",
	},
	"ButcherCleaver": {
		Name: "Cleaver",
		Text: "Basic Attacks deal area damage",
	},
	"ButcherHeroicAbilityButcherFurnaceBlast": {
		Name: "Furnace Blast",
		Text: "Deal damage around The Butcher after a delay",
	},
	"ButcherHeroicAbilityLambToTheSlaughter": {
		Name: "Lamb to the Slaughter",
		Text: "Chain an enemy Hero to a post",
	},
	"ButcherMasteryButchersBrandCraveFlesh": {
		Name: "Crave Flesh",
		Text: "Butcher's Brand grants Move Speed",
	},
	"ButcherMasteryButchersBrandExposeFlesh": {
		Name: "Expose Flesh",
		Text: "Expose FleshWhile a target is under the effects of Butcher's Brand, they are also Vulnerable, increasing all damage taken by 25%.",
	},
	"ButcherMasteryButchersBrandInsatiableBlade": {
		Name: "Insatiable Blade",
		Text: "Increases Move Speed while enemies are branded",
	},
	"ButcherMasteryFiresofHell": {
		Name: "Fires of Hell",
		Text: "Furnace Blast explodes twice",
	},
	"ButcherMasteryFreshMeatBloodFrenzy": {
		Name: "Blood Frenzy",
		Text: "Basic Attacks grant Attack and Move Speed",
	},
	"ButcherMasteryFreshMeatVictuals": {
		Name: "Victuals",
		Text: "Nearby Minion deaths heal",
	},
	"ButcherMasteryHamstringBrutalStrike": {
		Name: "Brutal Strike",
		Text: "Hamstring empowers Basic Attacks",
	},
	"ButcherMasteryHamstringCheapShot": {
		Name: "Cheap Shot",
		Text: "Increase Hamstring damage to disabled targets",
	},
	"ButcherMasteryHamstringChopMeat": {
		Name: "Chop Meat",
		Text: "Increases Hamstring damage to non-Heroes",
	},
	"ButcherMasteryHamstringCripplingSlam": {
		Name: "Crippling Slam",
		Text: "Hamstring slow no longer fades and lasts longer",
	},
	"ButcherMasteryHamstringFlailAxe": {
		Name: "Flail Axe",
		Text: "Increases Hamstring length",
	},
	"ButcherMasteryHamstringInvigoration": {
		Name: "Invigoration",
		Text: "Reduces Hamstring Mana cost, cooldown",
	},
	"ButcherMasteryRuthlessOnslaughtSavageCharge": {
		Name: "Savage Charge",
		Text: "Ruthless Onslaught deals extra damage",
	},
	"ButcherMasteryRuthlessOnslaughtUnrelentingPursuit": {
		Name: "Unrelenting Pursuit",
		Text: "Reduces Ruthless Onslaught cooldown on impact",
	},
	"ButcherMasterySlaughterhouse": {
		Name: "Slaughterhouse",
		Text: "Lamb to the Slaughter chains all enemy Heroes",
	},
	"ButcherMeatShield": {
		Name: "Meat Shield",
		Text: "Ruthless Onslaught grants Spell Armor",
	},
	"ButcherTalentEnraged": {
		Name: "Enraged",
		Text: "Gain Attack Speed and Armor at low Health",
	},
	"ChenAccumulatingFlame": {
		Name: "Accumulating Flame",
		Text: "Quest: Ignite Heroes to empower Breath of Fire",
	},
	"ChenAnotherRound": {
		Name: "Another Round",
		Text: "Keg Smash lowers cooldowns, has increased radius",
	},
	"ChenBolderFlavor": {
		Name: "Bolder Flavor",
		Text: "Fortifying Brew grants a large Shield immediately",
	},
	"ChenBrewmastersBalance": {
		Name: "Brewmaster's Balance",
		Text: "Gain bonuses based on current level of Brew",
	},
	"ChenDeadlyStrike": {
		Name: "Deadly Strike",
		Text: "Increase Flying Kick damage and remove Brew cost",
	},
	"ChenElementalConduit": {
		Name: "Elemental Conduit",
		Text: "Increase spirit Health and empower their abilities",
	},
	"ChenElusiveBrawler": {
		Name: "Elusive Brawler",
		Text: "Activate to evade Basic Attacks",
	},
	"ChenFlyingLeap": {
		Name: "Flying Leap",
		Text: "Increased Range",
	},
	"ChenFreshestIngredients": {
		Name: "Freshest Ingredients",
		Text: "Quest: Regen Globes increase Health Regen",
	},
	"ChenGroundingBrew": {
		Name: "Grounding Brew",
		Text: "Fortifying Brew grants Spell Armor",
	},
	"ChenHeroicAbilityStormEarthFire": {
		Name: "Storm, Earth, Fire",
		Text: "Split into three elemental spirits",
	},
	"ChenHeroicAbilityWanderingKeg": {
		Name: "Wandering Keg",
		Text: "Roll around in the barrel and knock away enemies",
	},
	"ChenKegToss": {
		Name: "Keg Toss",
		Text: "Quest: Hit Heroes to empower Keg Smash",
	},
	"ChenMasteryFortifyingBrewEnoughToShare": {
		Name: "Enough to Share",
		Text: "Fortifying Brew Shields allied Heroes",
	},
	"ChenMasteryKegSmashATouchOfHoney": {
		Name: "A Touch of Honey",
		Text: "Increase Keg Smash slow",
	},
	"ChenMasteryWanderingKegUntappedPotential": {
		Name: "Untapped Potential",
		Text: "Increases Wandering Keg Speed, Armor",
	},
	"ChenPressurePoint": {
		Name: "Pressure Point",
		Text: "Flying Kick slows enemies",
	},
	"ChenPurifyingBrew": {
		Name: "Purifying Brew",
		Text: "Empowers Fortifying Brew",
	},
	"ChenRefreshingElixir": {
		Name: "Refreshing Elixir",
		Text: "Healing effects greatly increased while drinking",
	},
	"ChenRingofFire": {
		Name: "Ring of Fire",
		Text: "Damage nearby enemies after using Breath of Fire",
	},
	"ChenStormstoutSecretRecipe": {
		Name: "Stormstout Secret Recipe",
		Text: "Basic Abilities heal Chen",
	},
	"ChenWitheringFlames": {
		Name: "Withering Flames",
		Text: "Setting Heroes on fire reduces their Spell Power",
	},
	"ChoCallousedHide": {
		Name: "Calloused Hide",
		Text: "Ogre Hide increases healing received",
	},
	"ChoConsumingFire": {
		Name: "Consuming Fire",
		Text: "Increased healing against Heroes",
	},
	"ChoCthunsGift": {
		Name: "C'Thun's Gift",
		Text: "Basic Attacks are now ranged, slow targets",
	},
	"ChoEnragedRegeneration": {
		Name: "Enraged Regeneration",
		Text: "Gall's Ogre Rage increases Health Regen",
	},
	"ChoFavorOfTheOldGods": {
		Name: "Favor of the Old Gods",
		Text: "Upheaval roots Heroes",
	},
	"ChoFirestarter": {
		Name: "Firestarter",
		Text: "Reduced Consuming Blaze cooldown",
	},
	"ChoFrenziedFists": {
		Name: "Frenzied Fists",
		Text: "Surging Fist increases Attack Speed",
	},
	"ChoFuelForTheFlame": {
		Name: "Fuel for the Flame",
		Text: "Quest: Kills empower Consuming Blaze",
	},
	"ChoHeroicAbilityHammerOfTwilight": {
		Name: "Hammer of Twilight",
		Text: "Activate to knockback and Stun enemies",
	},
	"ChoHeroicAbilityUpheaval": {
		Name: "Upheaval",
		Text: "Pull enemies in a large area",
	},
	"ChoHourofTwilight": {
		Name: "Hour of Twilight",
		Text: "Reduced death timer",
	},
	"ChoPowerSurge": {
		Name: "Power Surge",
		Text: "Surging Fist cooldown lowers against Heroes",
	},
	"ChoRollback": {
		Name: "Rollback",
		Text: "Rune Bomb returns",
	},
	"ChoRunedGauntlet": {
		Name: "Runed Gauntlet",
		Text: "Basic Attacks reduce Heroic cooldowns",
	},
	"ChoRunicFeedback": {
		Name: "Runic Feedback",
		Text: "Runic Blast reduces Rune Bomb cooldown",
	},
	"ChoSearedFlesh": {
		Name: "Seared Flesh",
		Text: "Consecutive Basic Attacks deal bonus damage",
	},
	"ChoSurgingDash": {
		Name: "Surging Dash",
		Text: "Unstoppable while channeling Surging Fist",
	},
	"ChoTalentMoltenBlock": {
		Name: "Molten Block",
		Text: "Activate to become Invulnerable",
	},
	"ChoTheWillofCho": {
		Name: "The Will of Cho",
		Text: "Takedowns increase Ogre Hide's Armor",
	},
	"ChoTwilightVeil": {
		Name: "Twilight Veil",
		Text: "Activate to increase Ogre Hide Armor",
	},
	"ChoUppercut": {
		Name: "Uppercut",
		Text: "Increases Surging Fist damage",
	},
	"ChromieBronzeTalons": {
		Name: "Bronze Talons",
		Text: "Increase Basic Attack range and damage",
	},
	"ChromieDragonsBreathDeepBreathing": {
		Name: "Deep Breathing",
		Text: "Quest: Hitting Heroes empowers Dragon's Breath",
	},
	"ChromieDragonsBreathDragonsEye": {
		Name: "Dragon’s Eye",
		Text: "Increase Dragon’s Breath center damage",
	},
	"ChromieDragonsBreathEnvelopingAssault": {
		Name: "Enveloping Assault",
		Text: "Increase the radius of Dragon’s Breath",
	},
	"ChromieDragonsBreathMobiusLoop": {
		Name: "Mobius Loop",
		Text: "Reduce Dragon’s Breath Mana, cooldown, damage",
	},
	"ChromieHereAndThere": {
		Name: "Here and There",
		Text: "Activate to swap positions with Sand Blast Echo",
	},
	"ChromieHeroicAbilitySlowingSands": {
		Name: "Slowing Sands",
		Text: "Create a Slowing area",
	},
	"ChromieHeroicAbilityTemporalLoop": {
		Name: "Temporal Loop",
		Text: "Return an enemy Hero to a previous position",
	},
	"ChromiePortBackToBaseByeBye": {
		Name: "Bye Bye!",
		Text: "Reduce Hearthstone cast time",
	},
	"ChromieQuantumOverdrive": {
		Name: "Quantum Overdrive",
		Text: "Activate to gain Spell Power",
	},
	"ChromieReachingThroughTime": {
		Name: "Reaching through Time",
		Text: "Increases Sand Blast and Dragon’s Breath range",
	},
	"ChromieSandBlastFastForward": {
		Name: "Fast Forward",
		Text: "Long range Sand Blasts reduce its cooldown",
	},
	"ChromieSandBlastMountingSand": {
		Name: "Mounting Sand",
		Text: "Increase Sand Blast consecutive bonus",
	},
	"ChromieSandBlastPiercingSands": {
		Name: "Piercing Sands",
		Text: "Sand Blast pierces",
	},
	"ChromieSandBlastShiftingSands": {
		Name: "Shifting Sands",
		Text: "Sand Blast grants Spell Power",
	},
	"ChromieSlowingSandsPocketOfTime": {
		Name: "Pocket of Time",
		Text: "Remove Slowing Sands Mana cost, increase Slow",
	},
	"ChromieTemporalLoopStuckInALoop": {
		Name: "Stuck in a Loop",
		Text: "Temporal Loop teleports twice",
	},
	"ChromieTimeOut": {
		Name: "Time Out",
		Text: "Activate to become Invulnerable",
	},
	"ChromieTimeTrapAndorhalAnomaly": {
		Name: "Andorhal Anomaly",
		Text: "Increase Time Trap charges",
	},
	"ChromieTimeTrapChronoSickness": {
		Name: "Chrono Sickness",
		Text: "Reduce Time Trap cooldown, Mana cost, adds Slow",
	},
	"ChromieTimeTrapTimelySurprise": {
		Name: "Timely Surprise",
		Text: "Time Trap resets Basic Abilities",
	},
	"ChromieTimewalkersPursuit": {
		Name: "Timewalker's Pursuit",
		Text: "Activate to reveal an areaQuest: Regen Globes increase Mana Regen",
	},
	"CrusaderBlessedHammer": {
		Name: "Blessed Hammer",
		Text: "Activate to create spinning hammers",
	},
	"CrusaderBlindedByTheLight": {
		Name: "Blinded By The Light",
		Text: "Activate to Shield nearby allies",
	},
	"CrusaderHeroicAbilityBlessedShield": {
		Name: "Blessed Shield",
		Text: "Throw a bouncing shield that Stuns",
	},
	"CrusaderHeroicAbilityFallingSword": {
		Name: "Falling Sword",
		Text: "Jump to area, damaging and Slowing enemies",
	},
	"CrusaderHeroicMasteryIndestructable": {
		Name: "Indestructible",
		Text: "Gain a Shield instead of dying",
	},
	"CrusaderHolyFury": {
		Name: "Holy Fury",
		Text: "Deals damage to nearby enemies",
	},
	"CrusaderMasteryBlessedShieldRadiatingFaith": {
		Name: "Radiating Faith",
		Text: "Increase Blessed Shield Stun duration, max targets",
	},
	"CrusaderMasteryCondemnConviction": {
		Name: "Conviction",
		Text: "Condemn increases Movement Speed",
	},
	"CrusaderMasteryCondemnEternalRetaliation": {
		Name: "Eternal Retaliation",
		Text: "Condemn cooldown reduced per enemy hit",
	},
	"CrusaderMasteryCondemnGravitationalPull": {
		Name: "Gravitational Pull",
		Text: "Gravitational Pull Range increased by 25%.",
	},
	"CrusaderMasteryCondemnKnightTakesPawn": {
		Name: "Knight Takes Pawn",
		Text: "Increases Condemn effect versus Minions and Mercenaries",
	},
	"CrusaderMasteryFallingSwordHeavensFury": {
		Name: "Heaven's Fury",
		Text: "Falling Sword bombards enemies with holy bolts",
	},
	"CrusaderMasteryIronSkinFanaticism": {
		Name: "Fanaticism",
		Text: "Damage taken during Iron Skin grants Move Speed",
	},
	"CrusaderMasteryIronSkinHoldYourGround": {
		Name: "Hold Your Ground",
		Text: "Increases Iron Skin Shield, reduced cooldown",
	},
	"CrusaderMasteryIronSkinReinforce": {
		Name: "Reinforce",
		Text: "Basic Abilities grant Physical Armor",
	},
	"CrusaderMasteryIronSkinTheCrusadeMarchesOn": {
		Name: "The Crusade Marches On",
		Text: "Abilities reduce the cooldown of Iron Skin",
	},
	"CrusaderMasteryLawsOfHope": {
		Name: "Laws of Hope",
		Text: "Increases Health Regen, activate to burst heal",
	},
	"CrusaderMasteryPunishRighteousSmash": {
		Name: "Righteous Smash",
		Text: "Punish restores Mana per enemy hit",
	},
	"CrusaderMasteryPunishRoar": {
		Name: "Roar",
		Text: "Increase Punish damage",
	},
	"CrusaderMasteryPunishSubdue": {
		Name: "Subdue",
		Text: "Quest: Increase Punish Slow",
	},
	"CrusaderMasteryShieldGlareHolyRenewal": {
		Name: "Holy Renewal",
		Text: "Shield Glare heals for each Hero hit",
	},
	"CrusaderMasteryShieldGlarePathOfTheCrusade": {
		Name: "Path Of The Crusade",
		Text: "Path of the CrusadeUsing Punish while Iron Skin is active refreshes the duration of Iron Skin.",
	},
	"CrusaderMasteryShieldGlareSinsExposed": {
		Name: "Sins Exposed",
		Text: "Shield Glare makes enemies take bonus damage",
	},
	"CrusaderZealousGlare": {
		Name: "Zealous Glare",
		Text: "Increase Shield Glare duration",
	},
	"DVaBigShot": {
		Name: "Big Shot",
		Text: "Damage enemies in a line as a Pilot",
	},
	"DVaBigShotPewPewPew": {
		Name: "Pew! Pew! Pew!",
		Text: "Big Shot fires multiple pulses",
	},
	"DVaBoostersComingThrough": {
		Name: "Coming Through",
		Text: "Increase Boosters knockback",
	},
	"DVaBoostersCrashCourse": {
		Name: "Crash Course",
		Text: "Quest: Reduce Boosters cooldown",
	},
	"DVaBoostersHitTheNitrous": {
		Name: "Hit the Nitrous",
		Text: "Boosters give extra immediate speed",
	},
	"DVaBoostersRushdown": {
		Name: "Rush-down",
		Text: "Reduced Boosters cooldown if used out of combat",
	},
	"DVaBunnyHop": {
		Name: "Bunny Hop",
		Text: "Become Unstoppable and Slow nearby Heroes",
	},
	"DVaBunnyHopStopAndPop": {
		Name: "Stop and Pop",
		Text: "Bunny Hop damage increased if stationary",
	},
	"DVaDefenseMatrixAggressionMatrix": {
		Name: "Aggression Matrix",
		Text: "Basic Attacks lower Defense Matrix cooldown",
	},
	"DVaDefenseMatrixDazerZone": {
		Name: "Dazer Zone",
		Text: "Defense Matrix Slows",
	},
	"DVaDefenseMatrixDivertingPower": {
		Name: "Diverting Power",
		Text: "Increase Defense Matrix size, Move Speed reduced",
	},
	"DVaDefenseMatrixFusionGenerator": {
		Name: "Fusion Generator",
		Text: "Defense Matrix adds Self-Destruct Charge",
	},
	"DVaDefenseMatrixGetThroughThis": {
		Name: "Get Through This!",
		Text: "Increase Defense Matrix duration",
	},
	"DVaMechAblativeArmor": {
		Name: "Ablative Armor",
		Text: "Take less damage from weak attacks",
	},
	"DVaMechEmergencyShielding": {
		Name: "Emergency Shielding",
		Text: "Mech gains a Shield instead of dying",
	},
	"DVaMechExpensivePlating": {
		Name: "Expensive Plating",
		Text: "Increase Mech Health and Call Mech cooldown",
	},
	"DVaMechProMoves": {
		Name: "Pro Moves",
		Text: "Taking damage increases Move Speed",
	},
	"DVaPilotCallMechMEKAfall": {
		Name: "MEKAfall",
		Text: "Call Mech is instant and targetable, damages enemies",
	},
	"DVaPilotConcussivePulse": {
		Name: "Concussive Pulse",
		Text: "Knockback enemies as a Pilot",
	},
	"DVaPilotGGWP": {
		Name: "GG, WP",
		Text: "Increased Pilot damage, Takedowns refresh Mech",
	},
	"DVaPilotNanoweaveSuit": {
		Name: "Nanoweave Suit",
		Text: "Faster Call Mech, ejection grants Armor",
	},
	"DVaPilotSuppressingFire": {
		Name: "Suppressing Fire",
		Text: "Pilot Basic Attacks Slow, range increased",
	},
	"DVaPilotTorpedoDash": {
		Name: "Torpedo Dash",
		Text: "Dash as a Pilot",
	},
	"DVaSelfDestructBringItOn": {
		Name: "Bring It On",
		Text: "Increase Self-Destruct Charge rate",
	},
	"DVaSelfDestructNuclearOption": {
		Name: "Nuclear Option",
		Text: "Increase Self-Destruct timer and damage",
	},
	"DehakaAdaptationChangeIsSurvivalTalent": {
		Name: "Change Is Survival",
		Text: "Enhanced Adaptation healing, cooldown",
	},
	"DehakaBrushstalkerApexPredator": {
		Name: "Apex Predator",
		Text: "Reduced Brushstalker cooldown, cast time",
	},
	"DehakaBrushstalkerFerociousStalker": {
		Name: "Ferocious Stalker",
		Text: "Brushstalker increases Dark Swarm damage",
	},
	"DehakaBurrowTalentLurkerStrain": {
		Name: "Lurker Strain",
		Text: "Burrow Stealths, knocks enemies away",
	},
	"DehakaDarkSwarmTalentEnduringSwarm": {
		Name: "Enduring Swarm",
		Text: "Dark Swarm grants Spell Armor",
	},
	"DehakaDarkSwarmTalentPrimalAggression": {
		Name: "Primal Aggression",
		Text: "Increases Dark Swarm damage to Minions, Mercs",
	},
	"DehakaDarkSwarmTalentSymbiosis": {
		Name: "Symbiosis",
		Text: "Reduces Dark Swarm cooldown against Heroes",
	},
	"DehakaDragTalentConstriction": {
		Name: "Constriction",
		Text: "Increases Drag duration",
	},
	"DehakaDragTalentElongatedTongue": {
		Name: "Elongated Tongue",
		Text: "Increases Drag range",
	},
	"DehakaDragTalentFeedingFrenzy": {
		Name: "Feeding Frenzy",
		Text: "Basic Attacks reduce Drag cooldown",
	},
	"DehakaDragTalentParalyzingEnzymes": {
		Name: "Paralyzing Enzymes",
		Text: "Drag slows",
	},
	"DehakaEssenceClaws": {
		Name: "Essence Claws",
		Text: "Basic Attacks slow, grant Essence",
	},
	"DehakaEssenceCollectionOneWhoCollectsTalent": {
		Name: "One-Who-Collects",
		Text: "Increases Essence from Minions",
	},
	"DehakaEssenceCollectionTalentEnhancedAgility": {
		Name: "Enhanced Agility",
		Text: "Quest: Essence increases Brushstalker Move Speed",
	},
	"DehakaEssenceCollectionTalentEssenceDevourer": {
		Name: "Essence Devourer",
		Text: "Quest: Regen Globes increase max Essence",
	},
	"DehakaEssenceCollectionTalentHeroStalker": {
		Name: "Hero Stalker",
		Text: "Increases Essence from Takedowns",
	},
	"DehakaEssenceCollectionTalentTissueRegeneration": {
		Name: "Tissue Regeneration",
		Text: "Quest: Essence increases Health Regen, max Essence",
	},
	"DehakaHeroicAbilityAdaptation": {
		Name: "Adaptation",
		Text: "Heal back damage taken",
	},
	"DehakaHeroicAbilityIsolation": {
		Name: "Isolation",
		Text: "Silence, Slow, and reduce vision of an enemy",
	},
	"DehakaIsolationTalentContagion": {
		Name: "Contagion",
		Text: "Isolation splashes to nearby Heroes",
	},
	"DehakaPrimalRage": {
		Name: "Primal Rage",
		Text: "Essence increases Attack Damage",
	},
	"DehakaPrimalSwarm": {
		Name: "Primal Swarm",
		Text: "Dark Swarm lowers enemies' Armor",
	},
	"DehakaSwiftPursuit": {
		Name: "Swift Pursuit",
		Text: "Increases Brushstalker bush move bonus duration",
	},
	"DehakaTunnelingClaws": {
		Name: "Tunneling Claws",
		Text: "Dehaka can move while Burrowed",
	},
	"DemonHunterCombatStyleHotPursuit": {
		Name: "Hot Pursuit",
		Text: "Max Hatred increases Movement Speed and Mana Regen",
	},
	"DemonHunterCombatStyleRancor": {
		Name: "Rancor",
		Text: "Hatred increases Attack Speed",
	},
	"DemonHunterCombatStyleTemperedByDiscipline": {
		Name: "Tempered by Discipline",
		Text: "Max Hatred grants healing on Basic Attacks",
	},
	"DemonHunterCreedoftheHunter": {
		Name: "Creed of the Hunter",
		Text: "Increases Attack SpeedQuest: Attacking Heroes empowers Hatred",
	},
	"DemonHunterDeathDealer": {
		Name: "Death Dealer",
		Text: "Increases Vault Basic Attack damage bonus",
	},
	"DemonHunterFarflightQuiver": {
		Name: "Farflight Quiver",
		Text: "Increases Basic Attack range",
	},
	"DemonHunterGloom": {
		Name: "Gloom",
		Text: "Consume Hatred to gain Spell Armor",
	},
	"DemonHunterHeroicAbilityRainofVengeance": {
		Name: "Rain of Vengeance",
		Text: "Area attack that stuns enemies",
	},
	"DemonHunterHeroicAbilityStrafe": {
		Name: "Strafe",
		Text: "Unleashes a flurry of arrows at enemies",
	},
	"DemonHunterManticore": {
		Name: "Manticore",
		Text: "Consecutive Basic Attacks against Heroes are empowered",
	},
	"DemonHunterMasteryArsenal": {
		Name: "Arsenal",
		Text: "Multishot launches grenades",
	},
	"DemonHunterMasteryCaltrops": {
		Name: "Caltrops",
		Text: "Drop Caltrops while Vaulting Quest: Maintain Hatred to empower Vault",
	},
	"DemonHunterMasteryCompositeArrowsMultishot": {
		Name: "Composite Arrows",
		Text: "Increases range of Multishot",
	},
	"DemonHunterMasteryDeathSiphon": {
		Name: "Death Siphon",
		Text: "Fires penetrating bolts and heals",
	},
	"DemonHunterMasteryFrostShot": {
		Name: "Frost Shot",
		Text: "Multishot slows enemies and has increased range",
	},
	"DemonHunterMasteryMonsterHunterHungeringArrow": {
		Name: "Monster Hunter",
		Text: "Reduces Hungering Arrow Mana cost and increases non-Hero damage",
	},
	"DemonHunterMasteryPuncturingArrow": {
		Name: "Puncturing Arrow",
		Text: "Quest: Hitting Heroes empowers Hungering Arrow",
	},
	"DemonHunterMasteryRepeatingArrowVault": {
		Name: "Repeating Arrow",
		Text: "Vault refreshes Hungering Arrow",
	},
	"DemonHunterMasterySiphoningArrow": {
		Name: "Siphoning Arrow",
		Text: "Hungering Arrow returns damage as Health",
	},
	"DemonHunterMasteryStormofVengeance": {
		Name: "Storm of Vengeance",
		Text: "Basic Attacks reduce Cooldown",
	},
	"DemonHunterMasteryTumble": {
		Name: "Tumble",
		Text: "Hold two charges of Vault",
	},
	"DemonHunterPunishment": {
		Name: "Punishment",
		Text: "Hatred increases Multishot damage",
	},
	"DemonHunterSeethingHatred": {
		Name: "Seething Hatred",
		Text: "Hatred stacks faster and lasts longer",
	},
	"DiabloBulwark": {
		Name: "Bulwark",
		Text: "Increases Shadow Charge Armor duration",
	},
	"DiabloDebilitatingFlames": {
		Name: "Debilitating Flames",
		Text: "Fire Stomp slows enemy Heroes",
	},
	"DiabloDevastatingCharge": {
		Name: "Devastating Charge",
		Text: "Quest: Increase Shadow Charge damage",
	},
	"DiabloDiabolicalMomentum": {
		Name: "Diabolical Momentum",
		Text: "Basic Attacks reduce Ability cooldowns",
	},
	"DiabloFearfulPresence": {
		Name: "Fearful Presence",
		Text: "Activate to reduce enemy damage",
	},
	"DiabloHellfire": {
		Name: "Hellfire",
		Text: "Increases Fire Stomp damage",
	},
	"DiabloHellgate": {
		Name: "Hellgate",
		Text: "Teleport and create an Apocalypse",
	},
	"DiabloHeroicAbilityApocalypse": {
		Name: "Apocalypse",
		Text: "Stun all enemy Heroes",
	},
	"DiabloHeroicAbilityLightningBreath": {
		Name: "Lightning Breath",
		Text: "Channeled spell that damages enemies",
	},
	"DiabloLifeLeech": {
		Name: "Life Leech",
		Text: "Basic Attacks against Heroes heal",
	},
	"DiabloMasteryDemonicStrength": {
		Name: "Demonic Strength",
		Text: "Overpower also slows",
	},
	"DiabloMasteryDevilsDueBlackSoulstone": {
		Name: "Devil's Due",
		Text: "Souls empower Regen Globes, Healing Fountains",
	},
	"DiabloMasteryDyingBreathApocalypse": {
		Name: "Dying Breath",
		Text: "Apocalypse triggers on death",
	},
	"DiabloMasteryFireDevilFireStomp": {
		Name: "Fire Devil",
		Text: "Fire Stomp grants a damage aura",
	},
	"DiabloMasteryFromTheShadowsShadowCharge": {
		Name: "From the Shadows",
		Text: "Increases Shadow Charge range",
	},
	"DiabloMasteryHellstormLightningBreath": {
		Name: "Hellstorm",
		Text: "Lightning Breath lasts and reaches longer",
	},
	"DiabloMasterySoulFeastBlackSoulstone": {
		Name: "Soul Feast",
		Text: "Black Soulstone increases Health Regen",
	},
	"DiabloSoulShield": {
		Name: "Soul Shield",
		Text: "Black Soulstone grants Spell Armor",
	},
	"DiabloSpeedDemon": {
		Name: "Speed Demon",
		Text: "Being Stunned, Rooted grants Movement Speed",
	},
	"DiabloTalentDominationOverpower": {
		Name: "Domination",
		Text: "Overpower resets cooldown of Shadow Charge",
	},
	"DiabloTalentLordOfTerror": {
		Name: "Lord of Terror",
		Text: "Steal Health from nearby enemy Heroes",
	},
	"DryadAbolishMagic": {
		Name: "Abolish Magic",
		Text: "Remove negative status effects from Lunara and the target",
	},
	"DryadBlossomSwell": {
		Name: "Blossom Swell",
		Text: "Increases Noxious Blossom radius",
	},
	"DryadBoundlessStrideTalent": {
		Name: "Boundless Stride",
		Text: "All Leaping Strike charges returned at once, can target allies",
	},
	"DryadChokingPollen": {
		Name: "Choking Pollen",
		Text: "Increases Noxious Blossom damage to enemies with Nature's Toxin",
	},
	"DryadCruelSpores": {
		Name: "Cruel Spores",
		Text: "Reduces Crippling Spores cooldown and Mana on Minions and Mercs",
	},
	"DryadDividingWisp": {
		Name: "Dividing Wisp",
		Text: "Relocating a Wisp creates a copy",
	},
	"DryadForestsWrath": {
		Name: "Forest's Wrath",
		Text: "Increases Thornwood Vine range and speed",
	},
	"DryadGallopingGait": {
		Name: "Galloping Gait",
		Text: "Activate to increase Movement Speed",
	},
	"DryadGreaterSpellShield": {
		Name: "Greater Spell Shield",
		Text: "Periodically gain temporary Spell Armor",
	},
	"DryadHeroicAbilityLeapingStrike": {
		Name: "Leaping Strike",
		Text: "Damages and leaps over an enemy",
	},
	"DryadHeroicAbilityThornwoodVine": {
		Name: "Thornwood Vine",
		Text: "Damages in a line",
	},
	"DryadInvigoratingSpores": {
		Name: "Invigorating Spores",
		Text: "Crippling Spores increases Basic Attack Speed",
	},
	"DryadLetThemWither": {
		Name: "Let Them Wither",
		Text: "Increases Crippling Spores slow",
	},
	"DryadNaturalPerspective": {
		Name: "Natural Perspective",
		Text: "Nature's Toxin reveals enemies",
	},
	"DryadNaturesCulling": {
		Name: "Nature's Culling",
		Text: "Increases Nature's Toxin damage to non-Heroes",
	},
	"DryadNimbleWisp": {
		Name: "Nimble Wisp",
		Text: "Increases Wisp Move Speed and vision",
	},
	"DryadPesteringBlossom": {
		Name: "Pestering Blossom",
		Text: "Increases Noxious Blossom range",
	},
	"DryadPhotosynthesis": {
		Name: "Photosynthesis",
		Text: "Crippling Spores returns Mana",
	},
	"DryadSiphoningToxin": {
		Name: "Siphoning Toxin",
		Text: "Nature's Toxin heals Lunara",
	},
	"DryadSkyboundWisp": {
		Name: "Skybound Wisp",
		Text: "Wisp sees over obstacles and reveals area on death",
	},
	"DryadSplinteredSpear": {
		Name: "Splintered Spear",
		Text: "Noxious Blossom causes Basic Attack to hit 4 targets",
	},
	"DryadStarWoodSpear": {
		Name: "Star Wood Spear",
		Text: "Crippling Spores increases Basic Attack range",
	},
	"DryadTimelostWisp": {
		Name: "Timelost Wisp",
		Text: "Wisp costs no Mana and reduces cooldown",
	},
	"DryadUnfairAdvantage": {
		Name: "Unfair Advantage",
		Text: "Increases Nature's Toxin damage to slowed Heroes",
	},
	"DryadWildVigor": {
		Name: "Wild Vigor",
		Text: "Crippling Spores increases Basic Attack damage",
	},
	"ETCBlockParty": {
		Name: "Block Party",
		Text: "Rockstar grants allies Physical Armor",
	},
	"ETCCombatStyleEchoPedal": {
		Name: "Echo Pedal",
		Text: "Abilities cause area damage",
	},
	"ETCCrowdSurfer": {
		Name: "Crowd Surfer",
		Text: "Powerslide can go over terrain",
	},
	"ETCHeroicAbilityMoshPit": {
		Name: "Mosh Pit",
		Text: "Nearby enemies must dance",
	},
	"ETCHeroicAbilityStageDive": {
		Name: "Stage Dive",
		Text: "Dive to a location and deal damage",
	},
	"ETCMasteryEncore": {
		Name: "Encore",
		Text: "Face Melt will knock enemies away a second time",
	},
	"ETCMasteryFaceSmelt": {
		Name: "Face Smelt",
		Text: "Face Melt slows enemies",
	},
	"ETCMasteryGuitarHero": {
		Name: "Guitar Hero",
		Text: "Basic Attacks heal during Guitar Solo",
	},
	"ETCMasteryGuitarSoloAggressiveShredding": {
		Name: "Aggressive Shredding",
		Text: "Basic Attacks reduce Guitar Solo cooldown",
	},
	"ETCMasteryHammeron": {
		Name: "Hammer-on",
		Text: "Abilities increase Basic Attack damage",
	},
	"ETCMasteryJustKeepRocking": {
		Name: "Just Keep Rockin'",
		Text: "Guitar Solo reduces disable durations",
	},
	"ETCMasteryMicCheck": {
		Name: "Mic Check",
		Text: "Reduces Face Melt cooldown when hitting multiple enemies",
	},
	"ETCMasteryProgRock": {
		Name: "Prog Rock",
		Text: "Quest: Regen Globes empower Guitar Solo",
	},
	"ETCMasteryRollingLikeaStone": {
		Name: "Rolling Like a Stone",
		Text: "Increased range on Powerslide",
	},
	"ETCMasteryShowStopper": {
		Name: "Show Stopper",
		Text: "Powerslide gives Armor",
	},
	"ETCMasterySpeedMetal": {
		Name: "Speed Metal",
		Text: "Move Speed is added to Rockstar ",
	},
	"FaerieDragonHardenedFocus": {
		Name: "Hardened Focus",
		Text: "Faster cooldowns near full Health",
	},
	"FaerieDragonHeroicAbilityBlinkHeal": {
		Name: "Blink Heal",
		Text: "Teleport to allies and heal them",
	},
	"FaerieDragonHeroicAbilityEmeraldWind": {
		Name: "Emerald Wind",
		Text: "Damages and pushes enemies in area",
	},
	"FaerieDragonMasteryArcanePrecision": {
		Name: "Arcane Precision",
		Text: "Increases Arcane Flare inner damage",
	},
	"FaerieDragonMasteryContinuousWinds": {
		Name: "Continuous Winds",
		Text: "Emerald Wind gains two additional novas",
	},
	"FaerieDragonMasteryCritterize": {
		Name: "Critterize",
		Text: "Polymorph lowers enemy's Armor",
	},
	"FaerieDragonMasteryPhaseShield": {
		Name: "Phase Shield",
		Text: "Grant Shields upon arrival",
	},
	"FaerieDragonMasteryShieldDust": {
		Name: "Shield Dust",
		Text: "Pixie Dust grants Physical Armor",
	},
	"FaerieDragonMasteryStickyFlare": {
		Name: "Sticky Flare",
		Text: "Arcane Flare slows enemies",
	},
	"FalstadHammerangBOOMerang": {
		Name: "BOOMerang",
		Text: "Activate Hammerang second time for more damage",
	},
	"FalstadHammerangGatheringStorm": {
		Name: "Gathering Storm",
		Text: "Quest: Hitting Heroes empowers Hammerang",
	},
	"FalstadHeroicAbilityHinterlandBlast": {
		Name: "Hinterland Blast",
		Text: "Long range damage beam",
	},
	"FalstadHeroicAbilityMightyGust": {
		Name: "Mighty Gust",
		Text: "Push away enemies and slow them",
	},
	"FalstadMasteryAerieGustsTailwind": {
		Name: "Aerie Gusts",
		Text: "Reduce Tailwind delay, increase Move Speed",
	},
	"FalstadMasteryAfterburner": {
		Name: "Afterburner",
		Text: "Barrel Roll increases Move Speed",
	},
	"FalstadMasteryBarrelRollFlowRider": {
		Name: "Flow Rider",
		Text: "Tailwind recharges Basic Abilities faster",
	},
	"FalstadMasteryBarrelRollFreeRoll": {
		Name: "Free Roll",
		Text: "Barrel Roll costs no Mana",
	},
	"FalstadMasteryCalloftheWildhammerHinterlandBlast": {
		Name: "Call of the Wildhammer",
		Text: "Increases Hinterland Blast range and damage",
	},
	"FalstadMasteryCripplingHammerHammerang": {
		Name: "Crippling Hammer",
		Text: "Increases Hammerang slow",
	},
	"FalstadMasteryDogFightBarrelRoll": {
		Name: "Dog Fight",
		Text: "Dog Fight The Barrel Roll Shield lasts 3 seconds longer.",
	},
	"FalstadMasteryFlightEpicMount": {
		Name: "Epic Mount",
		Text: "Reduces Flight cooldown and delay",
	},
	"FalstadMasteryFlyAwayFlight": {
		Name: "Fly Away!",
		Text: "Fly Away! The cooldown for Flight is reduced by  seconds.",
	},
	"FalstadMasteryHammerangPowerThrow": {
		Name: "Power Throw",
		Text: "Hammerang travels farther and slows longer",
	},
	"FalstadMasteryHammerangSecretWeapon": {
		Name: "Secret Weapon",
		Text: "Increases Hammerang range, Basic Attack damage",
	},
	"FalstadMasteryLightningRodChargedUp": {
		Name: "Charged Up",
		Text: "Increases Lightning Rod strikes and range",
	},
	"FalstadMasteryLightningRodLightningChain": {
		Name: "Lightning Chain",
		Text: "Lightning Chain Lightning Rod chains to up to 2 additional targets for 25% damage each. ",
	},
	"FalstadMasteryLightningRodStaticShield": {
		Name: "Static Shield",
		Text: "Gain Shields after every Lightning Rod strike",
	},
	"FalstadMasteryLightningRodThunderstrikes": {
		Name: "Thunderstrikes",
		Text: "Lightning Rod deals increasing damage",
	},
	"FalstadMasteryMightyGustWindTunnel": {
		Name: "Wind Tunnel",
		Text: "Mighty Gust constantly knocks back",
	},
	"FalstadMasteryUpdraftBarrelRoll": {
		Name: "Updraft",
		Text: "Increases Barrel Roll range and Shield",
	},
	"FalstadTalentHammerGains": {
		Name: "Hammer Gains",
		Text: "Basic Attacks heal",
	},
	"FalstadWingman": {
		Name: "Wingman",
		Text: "Kill Minions to bribe a Mercenary and empower Lightning Rod",
	},
	"GallBombsAway": {
		Name: "Bomb's Away",
		Text: "Quest: Empower Runic Blast",
	},
	"GallDeafeningBlast": {
		Name: "Deafening Blast",
		Text: "Runic Blast silences",
	},
	"GallDoubleBack": {
		Name: "Double Back",
		Text: "Re-activate Dread Orb to reverse third bounce",
	},
	"GallDoubleTrouble": {
		Name: "Double Trouble",
		Text: "Quest: Lower Shadowflame's cooldown",
	},
	"GallEdgeOfMadness": {
		Name: "Edge of Madness",
		Text: "Increase Shadowflame damage",
	},
	"GallEyeOfKilrogg": {
		Name: "Eye of Kilrogg",
		Text: "Place an eye that grants vision",
	},
	"GallHeroicAbilityShadowboltVolley": {
		Name: "Shadow Bolt Volley",
		Text: "Launch a barrage of Shadow Bolts",
	},
	"GallHeroicAbilityTwistingNether": {
		Name: "Twisting Nether",
		Text: "Slow and damage nearby enemies",
	},
	"GallKeepMoving": {
		Name: "Keep Moving!",
		Text: "Increased Shove Movespeed duration",
	},
	"GallLeadenOrb": {
		Name: "Leaden Orb",
		Text: "Dread Orb stuns",
	},
	"GallOgreRampage": {
		Name: "Ogre Rampage",
		Text: "Cho's Ogre Hide reduces cooldowns",
	},
	"GallPsychoticBreak": {
		Name: "Psychotic Break",
		Text: "Cast Abilities after death",
	},
	"GallRisingDread": {
		Name: "Rising Dread",
		Text: "Increase Dread Orb area and damage",
	},
	"GallRunicPersistence": {
		Name: "Runic Persistence",
		Text: "Runic Blast deals additional damage over time",
	},
	"GallSearingShadows": {
		Name: "Searing Shadows",
		Text: "Increased Shadowflame Hero damage",
	},
	"GallShadowboltVolleyShadowfury": {
		Name: "Shadowfury",
		Text: "Shadow Bolt Volley pierces",
	},
	"GallShadowsnare": {
		Name: "Shadowsnare",
		Text: "Shadowflame slows Heroes",
	},
	"GallShiftingNether": {
		Name: "Shifting Nether",
		Text: "Teleport and cast Twisting Nether",
	},
	"GallShove": {
		Name: "Shove",
		Text: "Nudge Cho a short distance",
	},
	"GallTaskmaster": {
		Name: "Taskmaster",
		Text: "Reduce Shove's cooldown",
	},
	"GallTheWillofGall": {
		Name: "The Will of Gall",
		Text: "Takedowns increase Ogre Rage's Damage Bonus",
	},
	"GallTwilightNova": {
		Name: "Twilight Nova",
		Text: "Dread Orb launches 2 extra bombs",
	},
	"GallWeSeeYou": {
		Name: "We See You!",
		Text: "Lower Eye of Kilrogg cooldown",
	},
	"GarroshArmorUpBodyCheck": {
		Name: "Body Check",
		Text: "Damage and Slow an enemy",
	},
	"GarroshArmorUpDoubleUp": {
		Name: "Double Up",
		Text: "Activate Armor Up to increase bonus",
	},
	"GarroshArmorUpInnerRage": {
		Name: "Inner Rage",
		Text: "Body Check gains an additional charge",
	},
	"GarroshBloodthirstBloodcraze": {
		Name: "Bloodcraze",
		Text: "Bloodthirst heals over time",
	},
	"GarroshBloodthirstInFortheKill": {
		Name: "In For the Kill",
		Text: "Bloodthirst kills reset cooldown, restore Mana",
	},
	"GarroshBloodthirstThirstforBattle": {
		Name: "Thirst for Battle",
		Text: "Basic Attacks reduce Bloodthirst cooldown",
	},
	"GarroshBodyCheckBruteForce": {
		Name: "Brute Force",
		Text: "Body Check reduces healing received",
	},
	"GarroshDecimateDeadlyCalm": {
		Name: "Deadly Calm",
		Text: "Decimate reduces damage dealt",
	},
	"GarroshGroundbreakerDefensiveMeasures": {
		Name: "Defensive Measures",
		Text: "Groundbreaker pulls grant Shield",
	},
	"GarroshGroundbreakerIntimidation": {
		Name: "Intimidation",
		Text: "Groundbreaker reduces Attack Speed",
	},
	"GarroshGroundbreakerMortalCombo": {
		Name: "Mortal Combo",
		Text: "Reduce Wrecking Ball cooldown on pulled targets",
	},
	"GarroshGroundbreakerRoughLanding": {
		Name: "Rough Landing",
		Text: "Groundbreaker Slows pulled Heroes",
	},
	"GarroshGroundbreakerWarbreaker": {
		Name: "Warbreaker",
		Text: "Quest: Pull Heroes to empower Groundbreaker",
	},
	"GarroshHeroicAbilityDecimate": {
		Name: "Decimate",
		Text: "Damage and Slow nearby enemies",
	},
	"GarroshHeroicAbilityWarlordsChallenge": {
		Name: "Warlord's Challenge",
		Text: "Force nearby Heroes to attack Garrosh",
	},
	"GarroshIndomitable": {
		Name: "Indomitable",
		Text: "Activate to become Unstoppable",
	},
	"GarroshOppressor": {
		Name: "Oppressor ",
		Text: "Basic Attacks reduce Spell Power",
	},
	"GarroshWarlordsChallengeDeathWish": {
		Name: "Death Wish",
		Text: "Takedowns reduce Warlord's Challenge cooldown",
	},
	"GarroshWreckingBallEarthshaker": {
		Name: "Earthshaker",
		Text: "Wrecking Ball Stuns enemies near impact",
	},
	"GarroshWreckingBallIntotheFray": {
		Name: "Into the Fray",
		Text: "Activate to throw an ally",
	},
	"GarroshWreckingBallTitanicMight": {
		Name: "Titanic Might",
		Text: "Wrecking Ball throws an additional enemy",
	},
	"GarroshWreckingBallUnrivaledStrength": {
		Name: "Unrivaled Strength",
		Text: "Increase Wrecking Ball's throw range, damage",
	},
	"GenericArcanePower": {
		Name: "Arcane Power",
		Text: "Increases Spell Power and restores Mana",
	},
	"GenericDampenMagic": {
		Name: "Dampen Magic",
		Text: "Periodically gain Spell Armor",
	},
	"GenericDoesNothing": {
		Name: "Placeholder Talent",
		Text: "This talent does nothing.",
	},
	"GenericTalentAmplifiedHealing": {
		Name: "Amplified Healing",
		Text: "All healing effects increased",
	},
	"GenericTalentBattleMomentum": {
		Name: "Battle Momentum",
		Text: "Basic Attacks decrease Ability cooldowns",
	},
	"GenericTalentBerserk": {
		Name: "Berserk",
		Text: "Activate to increase Attack and Movement Speed",
	},
	"GenericTalentBlock": {
		Name: "Block",
		Text: "Periodically gain Physical Armor",
	},
	"GenericTalentBloodForBlood": {
		Name: "Blood for Blood",
		Text: "Activate to steal Health from your target",
	},
	"GenericTalentBribe": {
		Name: "Bribe",
		Text: "Kill Minions to bribe a Mercenary",
	},
	"GenericTalentBurningRage": {
		Name: "Burning Rage",
		Text: "Deals damage to nearby enemies",
	},
	"GenericTalentCalldownMULE": {
		Name: "Calldown: MULE",
		Text: "Activate to heal and rearm Structures",
	},
	"GenericTalentClairvoyance": {
		Name: "Clairvoyance",
		Text: "Activate to reveal target area and nearby enemies",
	},
	"GenericTalentCleanse": {
		Name: "Cleanse",
		Text: "Makes target ally Unstoppable",
	},
	"GenericTalentConjurersPursuit": {
		Name: "Conjurer's Pursuit",
		Text: "Quest: Regen Globes increase Mana Regeneration",
	},
	"GenericTalentDemolitionist": {
		Name: "Demolitionist",
		Text: "Basic Attacks destroy ammo and deal extra damage to Structures",
	},
	"GenericTalentEnvenom": {
		Name: "Envenom",
		Text: "Activate to deal damage over time",
	},
	"GenericTalentFirstAid": {
		Name: "First Aid",
		Text: "Activate to heal over time",
	},
	"GenericTalentFlashoftheStorms": {
		Name: "Bolt of the Storm",
		Text: "Activate to teleport a short distance",
	},
	"GenericTalentFocusedAttack": {
		Name: "Focused Attack",
		Text: "Periodically empower Basic Attacks against Heroes",
	},
	"GenericTalentFollowThrough": {
		Name: "Follow Through",
		Text: "Abilities increase Basic Attack damage",
	},
	"GenericTalentFuryoftheStorm": {
		Name: "Fury of the Storm",
		Text: "Basic Attacks chain to non-Heroes",
	},
	"GenericTalentGatheringPower": {
		Name: "Gathering Power",
		Text: "Quest: Takedowns grant Spell Power",
	},
	"GenericTalentGiantKiller": {
		Name: "Giant Killer",
		Text: "Empowers your Basic Attacks against Heroes",
	},
	"GenericTalentHardenedShield": {
		Name: "Hardened Shield",
		Text: "Activate to gain massive Armor",
	},
	"GenericTalentHealingWard": {
		Name: "Healing Ward",
		Text: "Activate to place a ward which heals in an area",
	},
	"GenericTalentIceBlock": {
		Name: "Ice Block",
		Text: "Activate to become Invulnerable",
	},
	"GenericTalentImposingPresence": {
		Name: "Imposing Presence",
		Text: "Activate to slow enemy Basic Attacks and Move Speed",
	},
	"GenericTalentMercenaryLord": {
		Name: "Mercenary Lord",
		Text: "Empower nearby Mercenaries",
	},
	"GenericTalentMinionKiller": {
		Name: "Minion Killer",
		Text: "Minion KillerDeal 25% increased damage to Minions, Mercenaries, and Summons.",
	},
	"GenericTalentNexusBlades": {
		Name: "Nexus Blades",
		Text: "Basic Attacks deal more damage and Slow",
	},
	"GenericTalentNexusFrenzy": {
		Name: "Nexus Frenzy",
		Text: "Increases Attack Speed and Range",
	},
	"GenericTalentNullificationShield": {
		Name: "Nullification Shield",
		Text: "Activate to gain massive Spell Armor",
	},
	"GenericTalentOverdrive": {
		Name: "Overdrive",
		Text: "Activate to empower your Abilities for extra Mana",
	},
	"GenericTalentPromote": {
		Name: "Promote",
		Text: "Target Minion takes less and deals more damage",
	},
	"GenericTalentProtectiveShield": {
		Name: "Protective Shield",
		Text: "Activate to shield an allied Hero",
	},
	"GenericTalentRegenerationMaster": {
		Name: "Regeneration Master",
		Text: "Quest: Regen Globes increase Health Regen",
	},
	"GenericTalentRelentless": {
		Name: "Relentless",
		Text: "Halves duration of disables",
	},
	"GenericTalentRewind": {
		Name: "Rewind",
		Text: "Activate to reset cooldowns",
	},
	"GenericTalentScoutingDrone": {
		Name: "Scouting Drone",
		Text: "Reveals a large area, can be killed",
	},
	"GenericTalentSearingAttacks": {
		Name: "Searing Attacks",
		Text: "Activate to increase Basic Attack damage at the cost of Mana",
	},
	"GenericTalentSeasonedMarksman": {
		Name: "Seasoned Marksman",
		Text: "Quest: Kill enemies to empower Basic Attacks",
	},
	"GenericTalentShrinkRay": {
		Name: "Shrink Ray",
		Text: "Activate to reduce an enemy Hero's damage",
	},
	"GenericTalentSpellShield": {
		Name: "Spell Shield",
		Text: "Periodically gain temporary Spell Armor",
	},
	"GenericTalentSprint": {
		Name: "Sprint",
		Text: "Activate to increase Move Speed",
	},
	"GenericTalentStoneskin": {
		Name: "Stoneskin",
		Text: "Activate to gain a Shield",
	},
	"GenericTalentStormShield": {
		Name: "Storm Shield",
		Text: "Activate to grant Shields to all nearby allies",
	},
	"GenericTalentSwiftStorm": {
		Name: "Swift Storm",
		Text: "You are no longer dismounted from taking damage.Increases mount speed by 20%.",
	},
	"GenericTalentVigorousAssault": {
		Name: "Vigorous Assault",
		Text: "Basic Attacks heal you",
	},
	"GenericVigorousStrikePassive": {
		Name: "Vigorous Strike",
		Text: "Basic Attacks heal you",
	},
	"GenjiCyberAgilityAgileDismount": {
		Name: "Agile Dismount",
		Text: "Increase Cyber Agility range while Mounted",
	},
	"GenjiCyberAgilityCyberShield": {
		Name: "Cyber Shield",
		Text: "Cyber Agility grants Spell Armor",
	},
	"GenjiCyberAgilityDoubleJump": {
		Name: "Double Jump",
		Text: "Cyber Agility gains a second charge",
	},
	"GenjiCyberAgilityPathfinder": {
		Name: "Pathfinder",
		Text: "Jumping over terrain increases Move Speed",
	},
	"GenjiDeflectAugmentedGuard": {
		Name: "Augmented Guard",
		Text: "Gain a Shield after Deflect ends",
	},
	"GenjiDeflectDragonClaw": {
		Name: "Dragon Claw",
		Text: "Activate to damage nearby enemies",
	},
	"GenjiDeflectPerfectDefense": {
		Name: "Perfect Defense",
		Text: "Reduce Deflect cooldown",
	},
	"GenjiDeflectReflect": {
		Name: "Reflect",
		Text: "Increase Deflect damage",
	},
	"GenjiDeflectZanshin": {
		Name: "Zanshin",
		Text: "Quest: Deflect hits all nearby Heroes",
	},
	"GenjiDodge": {
		Name: "Dodge",
		Text: "Periodically dodge Basic Attacks",
	},
	"GenjiDragonbladeTheDragonBecomesMe": {
		Name: "The Dragon Becomes Me",
		Text: "Hitting enemy Heroes increases duration",
	},
	"GenjiHeroicDragonblade": {
		Name: "Dragonblade",
		Text: "Unleash the Dragonblade",
	},
	"GenjiHeroicXStrike": {
		Name: "X-Strike",
		Text: "Deal heavy damage in a cross shape",
	},
	"GenjiShurikenSharpenedStars": {
		Name: "Sharpened Stars",
		Text: "Shuriken pierce",
	},
	"GenjiShurikenShingan": {
		Name: "Shingan",
		Text: "Increase Shuriken single-target damage",
	},
	"GenjiShurikenShurikenMastery": {
		Name: "Shuriken Mastery",
		Text: "Quest: Increase Shuriken damage and restore charges",
	},
	"GenjiSwiftStrikeFinalCut": {
		Name: "Final Cut",
		Text: "Swift Strike deals extra damage after a delay",
	},
	"GenjiSwiftStrikeFlowLikeWater": {
		Name: "Flow Like Water",
		Text: "Hitting Heroes reduces Swift Strike cooldown",
	},
	"GenjiSwiftStrikeSteadyBlade": {
		Name: "Steady Blade",
		Text: "Swift Strike deals increased damage for each Hero hit",
	},
	"GenjiSwiftStrikeStrikeAtTheHeart": {
		Name: "Strike At The Heart",
		Text: "Extra Swift Strike damage at its end",
	},
	"GenjiSwiftStrikeSwiftAsTheWind": {
		Name: "Swift as the Wind",
		Text: "Swift Strike grants Move Speed",
	},
	"GenjiXStrikeLivingWeapon": {
		Name: "Living Weapon",
		Text: "Hero hits reduce X-Strike's cooldown",
	},
	"GiantKillerTychus": {
		Name: "Problem Solver",
		Text: "Empowers Tychus's Basic Attacks against Heroes",
	},
	"GreymaneCursedBulletGilneanRoulette": {
		Name: "Gilnean Roulette",
		Text: "Cursed Bullet pierces, lower cooldown",
	},
	"GreymaneDarkflightCyclicalNature": {
		Name: "Cyclical Nature",
		Text: "Hitting a Hero with Gilnean Cocktail reduces Darkflight's cooldown",
	},
	"GreymaneDarkflightDisengageRunningWild": {
		Name: "Running Wild",
		Text: "Increases Darkflight and Disengage range",
	},
	"GreymaneDarkflightThickSkin": {
		Name: "Thick Skin",
		Text: "Darkflight grants Physical Armor",
	},
	"GreymaneDisengageEyesInTheDark": {
		Name: "Eyes in the Dark",
		Text: "Disengage grants Stealth",
	},
	"GreymaneEagerWolf": {
		Name: "Eager Wolf",
		Text: "Increases Inner Beast Attack Speed",
	},
	"GreymaneGilneanCocktailDraughtOverflow": {
		Name: "Draught Overflow",
		Text: "Increases Gilnean Cocktail explosion length",
	},
	"GreymaneGilneanCocktailIncendiaryElixir": {
		Name: "Incendiary Elixir",
		Text: "Quest: Increases Gilnean Cocktail damage",
	},
	"GreymaneGilneanCocktailPerfectAim": {
		Name: "Perfect Aim",
		Text: "Increases Gilnean Cocktail's range and refunds Mana against Heroes",
	},
	"GreymaneGoForTheThroatUnleashed": {
		Name: "Unleashed",
		Text: "Go for the Throat gets more powerful free casts",
	},
	"GreymaneHeroicAbilityCursedBullet": {
		Name: "Cursed Bullet",
		Text: "Remove a portion of enemy's current Health",
	},
	"GreymaneHeroicAbilityGoForTheThroat": {
		Name: "Go for the Throat",
		Text: "Eviscerate a Hero, refresh cooldown if they die",
	},
	"GreymaneHuntersBlunderbuss": {
		Name: "Hunter's Blunderbuss",
		Text: "Human Basic Attacks splash",
	},
	"GreymaneInnerBeastInsatiable": {
		Name: "Insatiable",
		Text: "Inner Beast causes Basic Attacks to restore Mana",
	},
	"GreymaneInnerBeastOnTheProwl": {
		Name: "On the Prowl",
		Text: "Inner Beast increases Movement Speed",
	},
	"GreymaneInnerBeastViciousness": {
		Name: "Viciousness",
		Text: "Increases Inner Beast duration, Abilities refresh",
	},
	"GreymaneInnerBeastWolfheart": {
		Name: "Wolfheart",
		Text: "Basic Attacks lower Inner Beast cooldown",
	},
	"GreymaneRazorSwipeUnfetteredAssault": {
		Name: "Unfettered Assault",
		Text: "Increases Razor Swipe range",
	},
	"GreymaneRazorSwipeVisceralAttacks": {
		Name: "Visceral Attacks",
		Text: "Basic Attacks reduce Razor Swipe's cooldown",
	},
	"GreymaneToothAndClaw": {
		Name: "Tooth and Claw",
		Text: "Worgen form Basic Attacks cleave",
	},
	"GreymaneWizenedDuelist": {
		Name: "Wizened Duelist",
		Text: "Quest: Takedowns increase Basic Attack damage",
	},
	"GreymaneWorgenFormAlphaKiller": {
		Name: "Alpha Killer",
		Text: "Increases Worgen Basic Attacks against Heroes",
	},
	"GreymaneWorgenFormQuicksilverBullets": {
		Name: "Quicksilver Bullets",
		Text: "Increases Human Basic Attack range",
	},
	"GuldanChaoticEnergy": {
		Name: "Chaotic Energy",
		Text: "Quest: Gather Regen Globes to reduce Mana costs",
	},
	"GuldanConsumeSoul": {
		Name: "Consume Soul",
		Text: "Instantly kill an enemy Minion to heal",
	},
	"GuldanCorruptionCurseOfExhaustion": {
		Name: "Curse of Exhaustion",
		Text: "Corruption slows enemy Movement Speed",
	},
	"GuldanCorruptionDemonicSight": {
		Name: "Demonic Sight",
		Text: "Corruption reveals enemies",
	},
	"GuldanCorruptionEchoedCorruption": {
		Name: "Echoed Corruption",
		Text: "Quest: Hit  Heroes with Corruption to increase strikes",
	},
	"GuldanCorruptionRuinousAffliction": {
		Name: "Ruinous Affliction",
		Text: "Corruption deals bonus damage",
	},
	"GuldanDarkBargain": {
		Name: "Dark Bargain",
		Text: "Increase Health, respawn time",
	},
	"GuldanDemonicCircle": {
		Name: "Demonic Circle",
		Text: "Summon a Demonic Circle that Gul'dan can teleport to",
	},
	"GuldanDrainLifeDevourTheFrail": {
		Name: "Devour the Frail",
		Text: "Drain Life deals more damage to weakened enemies",
	},
	"GuldanDrainLifeGlyphOfDrainLife": {
		Name: "Glyph of Drain Life",
		Text: "Increases the cast range of Drain Life",
	},
	"GuldanDrainLifeHarvestLife": {
		Name: "Harvest Life",
		Text: "Increases healing from Drain Life",
	},
	"GuldanDrainLifeHealthFunnel": {
		Name: "Health Funnel",
		Text: "Killing enemies with Drain Life refreshes its cooldown",
	},
	"GuldanFelFlameBoundByShadow": {
		Name: "Bound by Shadow",
		Text: "Fel Flame hits reduce the cooldown of Corruption",
	},
	"GuldanFelFlameFelArmor": {
		Name: "Fel Armor",
		Text: "Fel Flame grants Spell Armor",
	},
	"GuldanFelFlamePursuitOfFlame": {
		Name: "Pursuit of Flame",
		Text: "Quest: Hit 40 Heroes to increase area",
	},
	"GuldanFelFlameRampantHellfire": {
		Name: "Rampant Hellfire",
		Text: "Increases Fel Flame's damage",
	},
	"GuldanHealthstone": {
		Name: "Healthstone",
		Text: "Activate to heal",
	},
	"GuldanHorrify": {
		Name: "Horrify",
		Text: "Deals damage and Fears Heroes",
	},
	"GuldanHorrifyHaunt": {
		Name: "Haunt",
		Text: "Increases Horrify duration, lower enemies' Armor",
	},
	"GuldanLifeTapDarknessWithin": {
		Name: "Darkness Within",
		Text: "Life Tap to gain Spell Power",
	},
	"GuldanLifeTapHungerforPower": {
		Name: "Hunger for Power",
		Text: "Gain Spell Power but reduce healing received",
	},
	"GuldanLifeTapImprovedLifeTap": {
		Name: "Improved Life Tap",
		Text: "Life Tap restores more Mana",
	},
	"GuldanRainOfDestruction": {
		Name: "Rain of Destruction",
		Text: "Summon a rain of meteors in an area",
	},
	"GuldanRainOfDestructionDeepImpact": {
		Name: "Deep Impact",
		Text: "Rain of Destruction slows enemy Movement Speed",
	},
	"HeroGenericExecutionerPassive": {
		Name: "Executioner",
		Text: "Attacking disabled Heroes increases damage",
	},
	"IllidanBladesOfAzzinoth": {
		Name: "Blades of Azzinoth",
		Text: "Activate to increase Basic Attack damage",
	},
	"IllidanCombatStyleHuntersOnslaught": {
		Name: "Hunter's Onslaught",
		Text: "Damage from Basic Abilities heals",
	},
	"IllidanCombatStyleThrillOfBattle": {
		Name: "Thrill of Battle",
		Text: "Activate to double cooldown reduction",
	},
	"IllidanElusiveStrike": {
		Name: "Elusive Strike",
		Text: "Sweeping Strikes reduces Evasion Cooldown",
	},
	"IllidanFieryBrand": {
		Name: "Fiery Brand",
		Text: "Consecutive Basic Attacks to Heroes are empowered",
	},
	"IllidanHeroicAbilityMetamorphosis": {
		Name: "Metamorphosis",
		Text: "Damage area and increase max Health",
	},
	"IllidanHeroicAbilityTheHunt": {
		Name: "The Hunt",
		Text: "Charge a target from a very long range",
	},
	"IllidanMasteryBatteredAssaultSweepingStrike": {
		Name: "Battered Assault",
		Text: "Increases Sweeping Strike damage bonus",
	},
	"IllidanMasteryDemonicFormMetamorphosis": {
		Name: "Demonic Form",
		Text: "Demon form lasts forever, grants Attack Speed",
	},
	"IllidanMasteryFriendOrFoeDive": {
		Name: "Friend or Foe",
		Text: "Increases Dive range and can Dive to allies",
	},
	"IllidanMasteryImmolationSweepingStrike": {
		Name: "Immolation",
		Text: "Sweeping Strike deals area damage",
	},
	"IllidanMasteryLungeDive": {
		Name: "Lunge",
		Text: "Increased Dive range",
	},
	"IllidanMasteryMarkedforDeathDive": {
		Name: "Marked for Death",
		Text: "Repeated Dives deal extra damage",
	},
	"IllidanMasteryNowhereToHideTheHunt": {
		Name: "Nowhere to Hide",
		Text: "Increases The Hunt range and reveal low Health Heroes",
	},
	"IllidanMasteryRapidChaseDive": {
		Name: "Rapid Chase",
		Text: "Dive increases Move Speed",
	},
	"IllidanMasterySecondSweepSweepingStrike": {
		Name: "Second Sweep",
		Text: "Can store 2 charges of Sweeping Strike",
	},
	"IllidanMasteryShadowShieldEvasion": {
		Name: "Shadow Shield",
		Text: "Evasion grants a Shield",
	},
	"IllidanMasterySixthSenseEvasion": {
		Name: "Sixth Sense",
		Text: "Evasion grants Spell Armor",
	},
	"IllidanMasteryUnboundSweepingStrike": {
		Name: "Unbound",
		Text: "Quest: Sweeping Strike can go over terrain",
	},
	"IllidanNimbleDefender": {
		Name: "Nimble Defender",
		Text: "Sweeping Strikes grants Armor",
	},
	"IllidanReflexiveBlock": {
		Name: "Reflexive Block",
		Text: "Dive grants Physical Armor",
	},
	"IllidanSweepingStrikesThirstingBlade": {
		Name: "Thirsting Blade",
		Text: "Sweeping Strikes grants more Lifesteal",
	},
	"IllidanUnendingHatredPassive": {
		Name: "Unending Hatred",
		Text: "Quest: Killing enemies grants Basic Attack damage",
	},
	"JainaBlizzardSnowstorm": {
		Name: "Snowstorm",
		Text: "Increase Blizzard radius",
	},
	"JainaBlizzardStormFront": {
		Name: "Storm Front",
		Text: "Increase Blizzard cast range",
	},
	"JainaConeOfColdIceFloes": {
		Name: "Ice Floes",
		Text: "Increase Cone of Cold area, reduce cooldown",
	},
	"JainaConeOfColdNorthernExposure": {
		Name: "Northern Exposure",
		Text: "Cone of Cold lowers Armor",
	},
	"JainaConeOfColdNumbingBlast": {
		Name: "Numbing Blast",
		Text: "Cone of Cold Roots Chilled targets",
	},
	"JainaFingersOfFrost": {
		Name: "Fingers of Frost",
		Text: "Quest: Regen Globes increase Mana Regeneration",
	},
	"JainaFrigidTransmission": {
		Name: "Ice Blink",
		Text: "Activate to teleport and Chill nearby enemies",
	},
	"JainaFrostbiteArcaneIntellect": {
		Name: "Arcane Intellect",
		Text: "Damage to Chilled targets returns Mana",
	},
	"JainaFrostbiteDeepChill": {
		Name: "Deep Chill",
		Text: "Chill Slow stacks",
	},
	"JainaFrostbiteFrostArmor": {
		Name: "Frost Armor",
		Text: "Chill attackers and gain Physical Armor",
	},
	"JainaFrostbiteFrostbitten": {
		Name: "Frostbitten",
		Text: "Increases damage bonus",
	},
	"JainaFrostbiteIceBarrier": {
		Name: "Ice Barrier",
		Text: "Damaging Chilled targets grants Shield",
	},
	"JainaFrostbiteLingeringChill": {
		Name: "Lingering Chill",
		Text: "Increase Chill duration",
	},
	"JainaFrostboltFrostShards": {
		Name: "Frost Shards",
		Text: "Frostbolt pierces first target",
	},
	"JainaFrostboltIceLance": {
		Name: "Ice Lance",
		Text: "Reduce Frostbolt cooldown on Chilled targets",
	},
	"JainaFrostboltWintersReach": {
		Name: "Winter's Reach",
		Text: "Increase Frostbolt range",
	},
	"JainaHeroicRingOfFrost": {
		Name: "Ring of Frost",
		Text: "Create a ring that damages and Roots enemies ",
	},
	"JainaHeroicSummonWaterElemental": {
		Name: "Summon Water Elemental",
		Text: "Summon a Water Elemental that Chills enemies",
	},
	"JainaIcefuryWand": {
		Name: "Icefury Wand",
		Text: "Increase Basic Attack damage",
	},
	"JainaIcyVeins": {
		Name: "Icy Veins",
		Text: "Reduce cooldowns and Mana costs",
	},
	"JainaImprovedIceBlock": {
		Name: "Improved Ice Block",
		Text: "Gain Invulnerability and Chill enemies",
	},
	"JainaRingOfFrostColdSnap": {
		Name: "Cold Snap",
		Text: "Ring of Frost center deals damage, Roots",
	},
	"JainaSummonWaterElementalWintermute": {
		Name: "Wintermute",
		Text: "Increase cast range and mimics Basic Abilities",
	},
	"JunkratBombPacks": {
		Name: "Bomb Packs",
		Text: "Make Bomb Packs for allies to pick up",
	},
	"JunkratConcussionMineBOOMPOW": {
		Name: "BOOM POW",
		Text: "Hitting Heroes reduces Concussion Mine cooldown",
	},
	"JunkratConcussionMineBoggedDown": {
		Name: "Bogged Down",
		Text: "Concussion Mine Slows",
	},
	"JunkratConcussionMineBonzerHits": {
		Name: "Bonzer Hits",
		Text: "Quest: Increase Concussion Mine Knockback, damage",
	},
	"JunkratConcussionMineRipperAir": {
		Name: "Ripper Air",
		Text: "Concussion Mine knocks Junkrat farther",
	},
	"JunkratFragLauncherBouncyBouncy": {
		Name: "Bouncy Bouncy",
		Text: "Frag Launcher grenades gain damage after ricochet",
	},
	"JunkratFragLauncherBurstFire": {
		Name: "Burst Fire",
		Text: "Frag Launcher burst fires",
	},
	"JunkratFragLauncherCannonball": {
		Name: "Cannonball!",
		Text: "Bigger Basic Attack, Frag Launcher explosions",
	},
	"JunkratFragLauncherEndlessNades": {
		Name: "Endless Nades",
		Text: "Hitting Heroes reduces Frag Launcher cooldown",
	},
	"JunkratFragLauncherExtraWoundTimers": {
		Name: "Extra-Wound Timers",
		Text: "Frag Launcher grenades last longer",
	},
	"JunkratFragLauncherPutSomeEnglishOnIt": {
		Name: "Put Some English On It",
		Text: "Increase Frag Launcher distance",
	},
	"JunkratFragLauncherTasteForExplosions": {
		Name: "Taste For Explosions",
		Text: "Quest: Increase Frag Launcher damage",
	},
	"JunkratFragLauncherTrickyShuffles": {
		Name: "Tricky Shuffles",
		Text: "Gain Move Speed while Frag Launcher replenishing",
	},
	"JunkratHeroicAbilityRIPTire": {
		Name: "RIP-Tire",
		Text: "Create a motorized bomb that explodes",
	},
	"JunkratHeroicRocketRide": {
		Name: "Rocket Ride",
		Text: "Launch into the air, damaging enemies upon landing",
	},
	"JunkratRIPTireExtraOomph": {
		Name: "Extra Oomph",
		Text: "Increase RIP-Tire knockback, reduce cooldown",
	},
	"JunkratRocketRidePuckishScamp": {
		Name: "Puckish Scamp",
		Text: "Reduce Rocket Ride cooldown",
	},
	"JunkratSpreadVolley": {
		Name: "Spread Volley",
		Text: "Activate to make Frag Launcher spread",
	},
	"JunkratSteelTrapBigAs": {
		Name: "Big As",
		Text: "Increase Steel Trap radius, damage",
	},
	"JunkratSteelTrapChatteringTeeth": {
		Name: "Chattering Teeth",
		Text: "Steel Traps chase Heroes",
	},
	"JunkratSteelTrapGottaTrapEmAll": {
		Name: "Gotta Trap 'Em All!",
		Text: "Quest: Increase Steel Trap count",
	},
	"JunkratSteelTrapStickyWicket": {
		Name: "Sticky Wicket",
		Text: "Steel Trap Slows, no longer Roots",
	},
	"KaelthasArcaneDynamo": {
		Name: "Arcane Dynamo",
		Text: "Using Basic Abilities increases Spell Power",
	},
	"KaelthasFlamestrikeBurnedFlesh": {
		Name: "Burned Flesh",
		Text: "Empower Flamestrike against multiple Heroes",
	},
	"KaelthasFlamestrikeConvection": {
		Name: "Convection",
		Text: "Quest: Flamestrike hits increase damage",
	},
	"KaelthasFlamestrikeFuryOfTheSunwell": {
		Name: "Fury of the Sunwell",
		Text: "Flamestrike explodes twice",
	},
	"KaelthasFlamestrikeManaTap": {
		Name: "Mana Tap",
		Text: "Verdant Spheres restores Mana",
	},
	"KaelthasGravityLapseEnergyRoil": {
		Name: "Energy Roil",
		Text: "Reduce Gravity Lapse cooldown on Heroes",
	},
	"KaelthasGravityLapseGravityThrow": {
		Name: "Gravity Throw",
		Text: "Increases Gravity Lapse stun duration",
	},
	"KaelthasGravityLapseNetherWind": {
		Name: "Nether Wind",
		Text: "Increases Gravity Lapse range and restores Mana",
	},
	"KaelthasGravityLapseTriOptimal": {
		Name: "Tri-Optimal",
		Text: "Gravity Lapse reduces Verdant Spheres' cooldown",
	},
	"KaelthasHeroicAbilityPhoenix": {
		Name: "Phoenix",
		Text: "Summon a Phoenix to burn enemies",
	},
	"KaelthasHeroicAbilityPyroblast": {
		Name: "Pyroblast",
		Text: "Blast an enemy with a big fireball",
	},
	"KaelthasLivingBombBackdraft": {
		Name: "Backdraft",
		Text: "Living Bomb explosion slows enemies",
	},
	"KaelthasLivingBombChainBomb": {
		Name: "Chain Bomb",
		Text: "Living Bomb explosion spreads Living Bomb",
	},
	"KaelthasLivingBombFissionBomb": {
		Name: "Fission Bomb",
		Text: "Increases Living Bomb explosion radius",
	},
	"KaelthasLivingBombPyromaniac": {
		Name: "Pyromaniac",
		Text: "Living Bomb reduces Basic Ability cooldowns",
	},
	"KaelthasLivingBombSunKingsFury": {
		Name: "Sun King's Fury",
		Text: "Living Bomb deals more damage after spreading",
	},
	"KaelthasManaAddict": {
		Name: "Mana Addict",
		Text: "Quest: Regen Globes grant Arcane Barrier",
	},
	"KaelthasMasterOfFlames": {
		Name: "Master of Flames",
		Text: "Living Bomb's spread from explosions can now also spread Living Bomb.",
	},
	"KaelthasMasteryFlamethrower": {
		Name: "Flamethrower",
		Text: "Increases Flamestrike's range",
	},
	"KaelthasMasterySunfireEnchantment": {
		Name: "Sunfire Enchantment",
		Text: "Special attack on Verdant Spheres activation",
	},
	"KaelthasPhoenixRebirth": {
		Name: "Rebirth",
		Text: "Increases Phoenix duration and may retarget",
	},
	"KaelthasPyroblastPresenceOfMind": {
		Name: "Presence Of Mind",
		Text: "Increases Pyroblast radius and lowers cooldown",
	},
	"KaelthasTwinSpheres": {
		Name: "Twin Spheres",
		Text: "Verdant Spheres gains a second charge",
	},
	"KaelthasVerdantSpheresFelInfusion": {
		Name: "Fel Infusion",
		Text: "Increases Spell Power, Verdant Spheres heals",
	},
	"KelThuzadAcceleratedDecay": {
		Name: "Accelerated Decay",
		Text: "Death and Decay damage increases",
	},
	"KelThuzadArcaneEchoes": {
		Name: "Arcane Echoes",
		Text: "Reduce Death and Decay cooldown",
	},
	"KelThuzadArchlichArmor": {
		Name: "Armor of the Archlich",
		Text: "Activate to gain Physical Armor, Slow enemies",
	},
	"KelThuzadBarbedChains": {
		Name: "Barbed Chains",
		Text: "Chains deals more damage, reduces Armor",
	},
	"KelThuzadBlightedFrost": {
		Name: "Blighted Frost",
		Text: "Increase Frost Nova center damage, duration",
	},
	"KelThuzadChainLink": {
		Name: "Chain-Link",
		Text: "Pulling Heroes together reduces Chains cooldown",
	},
	"KelThuzadChainsOfIce": {
		Name: "Chains of Ice",
		Text: "Chains Slows after pull",
	},
	"KelThuzadChillingTouch": {
		Name: "Chilling Touch",
		Text: "Basic Attacks periodically deal area Spell damage",
	},
	"KelThuzadDeathchill": {
		Name: "Deathchill",
		Text: "Frost Blast Slows, Takedowns release blasts",
	},
	"KelThuzadFrozenTomb": {
		Name: "Frost Blast",
		Text: "Root and damage enemies around a target",
	},
	"KelThuzadGlacialSpike": {
		Name: "Glacial Spike",
		Text: "Create a spike that explodes, deals damage",
	},
	"KelThuzadHungeringCold": {
		Name: "Hungering Cold",
		Text: "Frost Nova makes enemies take bonus damage",
	},
	"KelThuzadIcyGrasp": {
		Name: "Icy Grasp",
		Text: "Increase Frost Nova Slow duration",
	},
	"KelThuzadMasterOfTheColdDark": {
		Name: "Master of the Cold Dark",
		Text: "Kel'Thuzad becomes increasingly stronger as he disrupts enemies",
	},
	"KelThuzadMightOfTheScourge": {
		Name: "Might of the Scourge",
		Text: "Hitting Heroes resets Shadow Fissure cooldown",
	},
	"KelThuzadPhylactery": {
		Name: "Phylactery of Kel'Thuzad",
		Text: "Quest: Collect Regen Globes for instant respawn",
	},
	"KelThuzadPowerOfIcecrown": {
		Name: "Power of Icecrown",
		Text: "Disabling enemies grants Spell Power",
	},
	"KelThuzadShadowFissure": {
		Name: "Shadow Fissure",
		Text: "Deal high damage in an area",
	},
	"KelThuzadShiftingMalice": {
		Name: "Shifting Malice",
		Text: "Activate to dash and deal damage",
	},
	"KelThuzadStripShields": {
		Name: "Strip Shields",
		Text: "Chains grants Shield, damages Shields",
	},
	"KelThuzadTheDamnedReturn": {
		Name: "The Damned Return",
		Text: "Create Shades that mimic Death and Decay",
	},
	"KelThuzadThePlaguelands": {
		Name: "The Plaguelands",
		Text: "Increase Death and Decay duration, radius",
	},
	"KerriganAssimilationMastery": {
		Name: "Assimilation Mastery",
		Text: "Assimilation lasts longer and increases Health and Mana regeneration",
	},
	"KerriganCombatStyleDoubleStrike": {
		Name: "Double Strike",
		Text: "Abilities empower Basic Attacks",
	},
	"KerriganCombatStyleFuryoftheSwarm": {
		Name: "Fury of the Swarm",
		Text: "Basic Attacks cleave",
	},
	"KerriganCombatStyleLingeringEssence": {
		Name: "Lingering Essence",
		Text: "Increases Assimilation Shield duration",
	},
	"KerriganEssenceForEssence": {
		Name: "Essence for Essence",
		Text: "Activate to steal Health from the target",
	},
	"KerriganHeroicAbilityMaelstrom": {
		Name: "Maelstrom",
		Text: "Persistently damage nearby enemies",
	},
	"KerriganHeroicAbilitySummonUltralisk": {
		Name: "Summon Ultralisk",
		Text: "Summon an Ultralisk to chase a target",
	},
	"KerriganMasteryAdaptation": {
		Name: "Adaptation",
		Text: "Ravage can jump to allies",
	},
	"KerriganMasteryAggressiveDefense": {
		Name: "Aggressive Defense",
		Text: "Increases Assimilation Shields gained",
	},
	"KerriganMasteryCrushingSwarm": {
		Name: "Impaling Swarm",
		Text: "Impaling Blades spawns  Zerglings",
	},
	"KerriganMasteryEviscerate": {
		Name: "Eviscerate",
		Text: "Increases Ravage range",
	},
	"KerriganMasteryImpalingBladesBladeTorrent": {
		Name: "Blade Torrent",
		Text: "Increases Impaling Blades area",
	},
	"KerriganMasteryImpalingBladesSharpenedBlades": {
		Name: "Sharpened Blades",
		Text: "Impaling Blades deals more damage",
	},
	"KerriganMasteryMaelstromOmegastorm": {
		Name: "Omegastorm",
		Text: "Increased Maelstrom size and Shield generation",
	},
	"KerriganMasteryPrimalGraspEnergizingGrasp": {
		Name: "Energizing Grasp",
		Text: "Primal Grasp costs less Mana",
	},
	"KerriganMasteryPrimalGraspPsionicPulse": {
		Name: "Psionic Pulse",
		Text: "Damage nearby enemies",
	},
	"KerriganMasteryRavageCleanKill": {
		Name: "Clean Kill",
		Text: "Ravage refunds Mana on kill and increases damage",
	},
	"KerriganMasteryRavageSiphoningImpact": {
		Name: "Siphoning Impact",
		Text: "Ravage heals Kerrigan",
	},
	"KerriganMasteryTorrasqueSummonUltralisk": {
		Name: "Torrasque",
		Text: "Ultralisk spawns a new Ultralisk on death",
	},
	"KerriganPsionicShift": {
		Name: "Psionic Shift",
		Text: "Activate to teleport a short distance, deal damage",
	},
	"KerriganQueensRush": {
		Name: "Queen's Rush",
		Text: "Takedowns increase Move Speed",
	},
	"L90ETCMasteryDeathMetal": {
		Name: "Death Metal",
		Text: "Dying causes Mosh Pit",
	},
	"L90ETCMasteryFaceMeltLoudSpeakers": {
		Name: "Loud Speakers",
		Text: "Increases range and knockback of Face Melt",
	},
	"L90ETCMasteryFaceMeltPinballWizard": {
		Name: "Pinball Wizard",
		Text: "Increases Face Melt damage after Powerslide",
	},
	"L90ETCMasteryGuitarSoloGroupies": {
		Name: "Groupies",
		Text: "Guitar Solo heals nearby allied Heroes",
	},
	"L90ETCMasteryMoshPitTourBus": {
		Name: "Tour Bus",
		Text: "Makes Powerslide usable during Mosh Pit",
	},
	"L90ETCMasteryStageDiveCrowdPleaser": {
		Name: "Crowd Pleaser",
		Text: "Stage Dive impact area bigger, cooldown reduced",
	},
	"L90ETCMasteryStageDiveRockGod": {
		Name: "Rock God!",
		Text: "Rock God! Reduces the cast and air time of Stage Dive by 50%, and increases the damage dealt by %.",
	},
	"LeoricBurningDespairTalent": {
		Name: "Burning Despair",
		Text: "Deal damage to nearby enemies",
	},
	"LeoricDrainHopeCrushingHopeTalent": {
		Name: "Crushing Hope",
		Text: "Full duration Drain Hope deals bonus damage",
	},
	"LeoricDrainHopeHardenedBonesTalent": {
		Name: "Hardened Bones",
		Text: "Gain Armor during Drain Hope",
	},
	"LeoricHeroicAbilityEntomb": {
		Name: "Entomb",
		Text: "Wall in an area",
	},
	"LeoricHeroicAbilityMarchoftheBlackKing": {
		Name: "March of the Black King",
		Text: "March forward, healing Leoric, damaging enemies",
	},
	"LeoricMasteryBuriedAliveEntomb": {
		Name: "Buried Alive",
		Text: "Entomb damages and Silences",
	},
	"LeoricMasteryConsumeVitalitySkeletalSwing": {
		Name: "Consume Vitality",
		Text: "Skeletal Swing heals",
	},
	"LeoricMasteryDeathMarchMarchoftheBlackKing": {
		Name: "Death March",
		Text: "March of the Black King applies Drain Hope",
	},
	"LeoricMasteryDrainMomentumDrainHope": {
		Name: "Drain Momentum",
		Text: "Drain Hope grants Move Speed",
	},
	"LeoricMasteryFealtyUntoDeathUndying": {
		Name: "Fealty Unto Death",
		Text: "Gain Health and Mana when Minions die",
	},
	"LeoricMasteryGhastlyReachSkeletalSwing": {
		Name: "Ghastly Reach",
		Text: "Increase Skeletal Swing range",
	},
	"LeoricMasteryHopelessnessDrainHope": {
		Name: "Hopelessness",
		Text: "Increase Drain Hope range",
	},
	"LeoricMasteryLingeringApparitionWraithWalk": {
		Name: "Lingering Apparition",
		Text: "Wraith Walk lasts longer",
	},
	"LeoricMasteryOsseinRenewal": {
		Name: "Ossein Renewal",
		Text: "Activate to heal",
	},
	"LeoricMasteryParalyzingRageSkeletalSwing": {
		Name: "Paralyzing Rage",
		Text: "Increase Skeletal Swing Slow",
	},
	"LeoricMasteryRoyalFocusWraithWalk": {
		Name: "Royal Focus",
		Text: "Empower Skeletal Swing and Wraith Walk",
	},
	"LeoricMasterySpectralLeech": {
		Name: "Spectral Leech",
		Text: "Basic Attacks deal extra damage and heal",
	},
	"LeoricMasteryUnyieldingDespairDrainHope": {
		Name: "Unyielding Despair",
		Text: "Reduce Drain Hope cooldown",
	},
	"LeoricMasteryWillingVesselDrainHope": {
		Name: "Willing Vessel",
		Text: "Increase Drain Hope healing",
	},
	"LeoricMithrilMaceTalent": {
		Name: "Mithril Mace",
		Text: "Quest: Empower Wrath Of The Bone King",
	},
	"LeoricShroudoftheDeadKingTalent": {
		Name: "Shroud of the Dead King",
		Text: "Activate to become Protected",
	},
	"LeoricSkeletalSwingKneelPeasantsTalent": {
		Name: "Kneel, Peasants!",
		Text: "Increase Skeletal Swing damage to non-Heroes",
	},
	"LeoricWraithWalkOminousWraithTalent": {
		Name: "Ominous Wraith",
		Text: "Empower Wraith Walk",
	},
	"LiLiHeroicAbilityJugof1000Cups": {
		Name: "Jug of 1,000 Cups",
		Text: "Rapidly heal nearby allies",
	},
	"LiLiHeroicAbilityWaterDragon": {
		Name: "Water Dragon",
		Text: "Damages and heavily slows nearest enemy Hero",
	},
	"LiLiMasteryBlindingWindGaleForce": {
		Name: "Gale Force",
		Text: "Increases Blinding Wind damage",
	},
	"LiLiMasteryBlindingWindHinderingWinds": {
		Name: "Hindering Winds",
		Text: "Blinding Wind slows targets",
	},
	"LiLiMasteryBlindingWindLingeringBlind": {
		Name: "Lingering Blind",
		Text: "Increases the duration of Blinding Wind",
	},
	"LiLiMasteryBlindingWindMassVortex": {
		Name: "Mass Vortex",
		Text: "Blinding Wind hits more targets",
	},
	"LiLiMasteryBlindingWindSurgingWinds": {
		Name: "Surging Winds",
		Text: "Gain Spell Power from Blinding Wind",
	},
	"LiLiMasteryCloudSerpentBringerofGifts": {
		Name: "Bringer of Gifts",
		Text: "Cloud Serpent also heals the target for  Health and  Mana.",
	},
	"LiLiMasteryCloudSerpentLightningSerpent": {
		Name: "Lightning Serpent",
		Text: "Cloud Serpent's attacks bounce",
	},
	"LiLiMasteryCloudSerpentMendingSerpent": {
		Name: "Mending Serpent",
		Text: "Cloud Serpent heals each attack",
	},
	"LiLiMasteryCloudSerpentSerpentSidekick": {
		Name: "Serpent Sidekick",
		Text: "Also gives Li Li a Cloud Serpent when cast on allies",
	},
	"LiLiMasteryCloudSerpentTimelessCreature": {
		Name: "Timeless Creature",
		Text: "Increases Cloud Serpent duration",
	},
	"LiLiMasteryFastFeetElusiveFeet": {
		Name: "Elusive Feet",
		Text: "Fast Feet grants Physical Armor",
	},
	"LiLiMasteryFastFeetSafetySprint": {
		Name: "Safety Sprint",
		Text: "Increases Fast Feet Move Speed",
	},
	"LiLiMasteryHealingBrewPitchPerfect": {
		Name: "Pitch Perfect",
		Text: "Repeated Healing Brews cost less Mana",
	},
	"LiLiMasteryHealingBrewProToss": {
		Name: "Pro Toss",
		Text: "Increases Healing Brew range",
	},
	"LiLiMasteryHealingBrewTheGoodStuff": {
		Name: "The Good Stuff",
		Text: "Healing Brew also heals over time",
	},
	"LiLiMasteryHealingBrewTwoForOne": {
		Name: "Two For One",
		Text: "Healing Brew heals twice but longer cooldown",
	},
	"LiLiMasteryJugof1000CupsJugof1000000Cups": {
		Name: "Jug of 1,000,000 Cups",
		Text: "Throw out two cups at a time",
	},
	"LiLiMasteryKungFuHustle": {
		Name: "Kung Fu Hustle",
		Text: "Cooldowns recharge faster when damaged",
	},
	"LiLiMasteryMagicalEssence": {
		Name: "Magical Essence",
		Text: "Magical EssenceIncreases the range of Healing Brew, Cloud Serpent, and Blinding Wind by 25%.",
	},
	"LiLiMasteryShakeItOff": {
		Name: "Shake It Off",
		Text: "Being Stunned or Rooted grants Armor",
	},
	"LiLiMasteryWaterDragonDoubleDragon": {
		Name: "Double Dragon",
		Text: "Summons two dragons",
	},
	"LostVikingsGoGoGo64KBMarathon": {
		Name: "64 KB Marathon",
		Text: "Run faster and break slows and roots",
	},
	"LostVikingsHeroicAbilityLongboatRaid": {
		Name: "Longboat Raid!",
		Text: "Load into a Longboat to attack enemies",
	},
	"LostVikingsHeroicAbilityPlayAgain": {
		Name: "Play Again!",
		Text: "Revive and summon Vikings",
	},
	"LostVikingsHeroicAbilityTheSequel": {
		Name: "The Sequel!",
		Text: "Reduces death timer for all Vikings",
	},
	"LostVikingsMasteryBaleogTheFierce": {
		Name: "Baleog the Fierce",
		Text: "Baleog's attacks increase his Attack Speed",
	},
	"LostVikingsMasteryCheckpointReached": {
		Name: "Checkpoint Reached",
		Text: "Additional revive and heal",
	},
	"LostVikingsMasteryErikTheSwift": {
		Name: "Erik the Swift",
		Text: "Increases Erik's Move Speed and heal while running",
	},
	"LostVikingsMasteryExplosiveAttacks": {
		Name: "Explosive Attacks",
		Text: "Increases Baleog splash to non-Heroic enemies",
	},
	"LostVikingsMasteryHunkaBurningOlaf": {
		Name: "Hunka' Burning Olaf",
		Text: "Olaf deals damage to nearby enemies",
	},
	"LostVikingsMasteryImpatienceIsAVirtue": {
		Name: "Impatience Is a Virtue",
		Text: "Vikings' attacks reduce cooldowns of all Vikings",
	},
	"LostVikingsMasteryItsASabotage": {
		Name: "It's a Sabotage!",
		Text: "Erik's attacks stronger against Structures",
	},
	"LostVikingsMasteryJump": {
		Name: "Jump!",
		Text: "Grants temporary Invulnerability",
	},
	"LostVikingsMasteryLargeAndInCharge": {
		Name: "Large and In Charge",
		Text: "Charge stuns enemies",
	},
	"LostVikingsMasteryNordicAttackSquad": {
		Name: "Nordic Attack Squad",
		Text: "Increases Viking Basic Attacks to Heroes",
	},
	"LostVikingsMasteryNorseForce": {
		Name: "Norse Force!",
		Text: "All Vikings gain a Shield",
	},
	"LostVikingsMasteryOlafTheStout": {
		Name: "Olaf the Stout",
		Text: "Olaf gains Physical Armor",
	},
	"LostVikingsMasteryPainDontHurt": {
		Name: "Pain Don't Hurt",
		Text: "Baleog's attacks heal himself",
	},
	"LostVikingsMasterySpinToWin": {
		Name: "Spin To Win!",
		Text: "Deals damage around each Viking",
	},
	"LostVikingsMasterySpyGames": {
		Name: "Spy Games",
		Text: "Erik gains Stealth and increases vision if stationary",
	},
	"LostVikingsMasteryTheSequel": {
		Name: "The Sequel!",
		Text: "Reduces death timer for all Vikings",
	},
	"LostVikingsMasteryVikingBribery": {
		Name: "Viking Bribery",
		Text: "Kill Minions to bribe a Mercenary",
	},
	"LostVikingsMasteryWereOnABoat": {
		Name: "Ragnarok 'n' Roll!",
		Text: "Increased combat prowess",
	},
	"LostVikingsTalentFuryoftheStorm": {
		Name: "Fury of the Storm",
		Text: "Basic Attacks chain to non-Heroes",
	},
	"LucioAmpItUpBringItTogether": {
		Name: "Bring it Together",
		Text: "Healing Boost increased near allies",
	},
	"LucioAmpItUpMaximumTempoQuest": {
		Name: "Maximum Tempo",
		Text: "Quest: Hero kills increase Amp It Up speed",
	},
	"LucioAmpItUpRejuvenescencia": {
		Name: "Rejuvenescência",
		Text: "Healing Boost also heals a percentage of life",
	},
	"LucioAmpItUpSonicAmplifier": {
		Name: "Sonic Amplifier",
		Text: "Amp It Up increases Crossfade radius",
	},
	"LucioAmpItUpSynaesthesiaAuditiva": {
		Name: "Synaesthesia Auditiva",
		Text: "Casting Amp It Up removes Stuns, Slows, and Roots",
	},
	"LucioAmpItUpUpTheFrequency": {
		Name: "Up the Frequency",
		Text: "Basic Attacks lower Amp It Up's cooldown",
	},
	"LucioAmptItUpBonusTrack": {
		Name: "Bonus Track",
		Text: "Using Crossfade refreshes Amp It Up",
	},
	"LucioBackInTheMix": {
		Name: "Back in the Mix",
		Text: "Heal self on Stun, Silence, or Time Stop",
	},
	"LucioBoombox": {
		Name: "Boombox",
		Text: "Drop a boombox that plays Crossfade",
	},
	"LucioCrossfadeBeatMixing": {
		Name: "Beat Mixing",
		Text: "Gain a Shield after switching Crossfade tracks",
	},
	"LucioCrossfadePartyMixQuest": {
		Name: "Party Mix",
		Text: "Quest: Playing Crossfade for Allies increases range",
	},
	"LucioCrossfadeSpeedBoostWeMoveTogether": {
		Name: "We Move Together",
		Text: "Increase Move Speed of Crossfade Speed Boost",
	},
	"LucioReverseAmp": {
		Name: "Reverse Amp",
		Text: "Cause Crossfade to damage, slow enemies",
	},
	"LucioReverseAmpNonstopRemix": {
		Name: "Nonstop Remix",
		Text: "Reverse Amp lasts forever near multiple enemies",
	},
	"LucioSoundBarrier": {
		Name: "Sound Barrier",
		Text: "Shield nearby allies",
	},
	"LucioSoundBarrierBossaNova": {
		Name: "Bossa Nova",
		Text: "Reduce Sound Barrier's cooldown, duration",
	},
	"LucioSoundwaveChaseTheBassQuest": {
		Name: "Chase the Bass",
		Text: "Quest: Enhance Soundwave's arc, range",
	},
	"LucioSoundwaveHeliotropics": {
		Name: "Heliotropics",
		Text: "Soundwave Blinds enemies",
	},
	"LucioSoundwaveOffTheWall": {
		Name: "Off the Wall",
		Text: "Soundwave's cooldown reduced during Wall Ride",
	},
	"LucioSoundwaveSubwoofer": {
		Name: "Subwoofer",
		Text: "Soundwave pushes close enemies further",
	},
	"LucioWallRideAccelerando": {
		Name: "Accelerando",
		Text: "Wall Ride gradually increases Move Speed",
	},
	"LucioWallRideCantStopWontStop": {
		Name: "Can't Stop, Won't Stop",
		Text: "Immune to Slow and Root effects during Wall Ride",
	},
	"LucioWallRideHardStyle": {
		Name: "Hard Style",
		Text: "Gain Armor during Wall Ride",
	},
	"LucioWallRideSlip": {
		Name: "Slip",
		Text: "Nearby enemy Heroes increase Wall Ride speed",
	},
	"MalfurionCombatStyleElunesGrace": {
		Name: "Elune's Grace",
		Text: "Increases Basic Ability ranges",
	},
	"MalfurionCombatStyleShandosClarity": {
		Name: "Shan'do's Clarity",
		Text: "Reduces Innervate cooldown",
	},
	"MalfurionHeroicAbilityTranquility": {
		Name: "Tranquility",
		Text: "Heals nearby allied Heroes over time",
	},
	"MalfurionHeroicAbilityTwilightDream": {
		Name: "Twilight Dream",
		Text: "Silence and damage nearby enemies",
	},
	"MalfurionMasteryAstralCommunion": {
		Name: "Astral Communion",
		Text: "Activate to teleport and cast Twilight Dream",
	},
	"MalfurionMasteryFullMoonfire": {
		Name: "Full Moonfire",
		Text: "Increases Moonfire area and reduces Mana cost",
	},
	"MalfurionMasteryHinderingMoonfire": {
		Name: "Hindering Moonfire",
		Text: "Moonfire Slows Heroes",
	},
	"MalfurionMasteryLifeSeed": {
		Name: "Life Seed",
		Text: "Nearby Heroes receive Regrowth periodically",
	},
	"MalfurionMasteryLunarShower": {
		Name: "Lunar Shower",
		Text: "Successive Moonfires are more powerful",
	},
	"MalfurionMasteryMoonburn": {
		Name: "Moonburn",
		Text: "Quest: Increases Moonfire damage to Minions",
	},
	"MalfurionMasterySerenity": {
		Name: "Serenity",
		Text: "Increases Tranquility healing and grants Mana",
	},
	"MalfurionMasteryStranglingVinesEntanglingRoots": {
		Name: "Strangling Vines",
		Text: "Entangling Roots reduces enemy healing",
	},
	"MalfurionMasteryTenaciousRootsEntanglingRoots": {
		Name: "Tenacious Roots",
		Text: "Increases Entangling Root area and duration",
	},
	"MalfurionMasteryVengefulRoots": {
		Name: "Vengeful Roots",
		Text: "Quest: Entangling Roots spawns a Treant",
	},
	"MalfurionRevitalizeInnervateTalent": {
		Name: "Revitalize",
		Text: "Innervate also benefits Malfurion",
	},
	"MalfurionYserasGift": {
		Name: "Ysera's Gift",
		Text: "Increase Regrowth heal at high Health",
	},
	"MalthaelAngelOfDeath": {
		Name: "Angel of Death",
		Text: "Last Rites heals",
	},
	"MalthaelBlackHarvest": {
		Name: "Black Harvest",
		Text: "Quest: Increase Reaper's Mark duration",
	},
	"MalthaelColdHand": {
		Name: "Cold Hand",
		Text: "Soul Rip Slows",
	},
	"MalthaelDeathsReach": {
		Name: "Death's Reach",
		Text: "Increase Wraith Strike range",
	},
	"MalthaelDieAlone": {
		Name: "Die Alone",
		Text: "Increase Soul Rip single-target damage",
	},
	"MalthaelEtherealExistence": {
		Name: "Ethereal Existence",
		Text: "Marked targets grant Physical Armor",
	},
	"MalthaelFearTheReaper": {
		Name: "Fear the Reaper",
		Text: "Activate to increase Movement Speed",
	},
	"MalthaelFinalCurtain": {
		Name: "Final Curtain",
		Text: "Death Shroud leaves a trail",
	},
	"MalthaelInevitableEnd": {
		Name: "Inevitable End",
		Text: "Activate to become Unstoppable, remove Marks",
	},
	"MalthaelLastRites": {
		Name: "Last Rites",
		Text: "Quest: Damage a Hero based on their missing Health",
	},
	"MalthaelMassacre": {
		Name: "Massacre",
		Text: "Wraith Strike hits an area",
	},
	"MalthaelMementoMori": {
		Name: "Memento Mori",
		Text: "Reaper's Mark deals increased damage over time",
	},
	"MalthaelMortality": {
		Name: "Mortality",
		Text: "Wraith Strike deals bonus damage to Heroes",
	},
	"MalthaelNoOneCanStopDeath": {
		Name: "No One Can Stop Death",
		Text: "Activate to respawn",
	},
	"MalthaelOnAPaleHorse": {
		Name: "On a Pale Horse",
		Text: "Increase mounted Movement Speed",
	},
	"MalthaelReaperOfSouls": {
		Name: "Reaper of Souls",
		Text: "Hero Takedowns extend Tormented Souls duration",
	},
	"MalthaelShroudOfWisdom": {
		Name: "Shroud of Wisdom",
		Text: "Activate to gain Spell Armor",
	},
	"MalthaelSoulCollector": {
		Name: "Soul Collector",
		Text: "Increase Soul Rip range, reduce cooldown",
	},
	"MalthaelSoulSiphon": {
		Name: "Soul Siphon",
		Text: "Increase Soul Rip healing",
	},
	"MalthaelThrowingShade": {
		Name: "Throwing Shade",
		Text: "Quest: Increase Death Shroud range",
	},
	"MalthaelTormentedSouls": {
		Name: "Tormented Souls",
		Text: "Apply Reaper's Mark to nearby enemies",
	},
	"MalthaelTouchOfDeath": {
		Name: "Touch of Death",
		Text: "Activate to reduce enemy healing",
	},
	"MedicBlastShield": {
		Name: "Blast Shield",
		Text: "Displacement Grenade grants Energy, Shield",
	},
	"MedicCaduceusFeedback": {
		Name: "Caduceus Feedback",
		Text: "Basic Attacks generate Energy",
	},
	"MedicCaduceusReactor2dot0": {
		Name: "Caduceus Reactor 2.0",
		Text: "Increase Caduceus Reactor healing",
	},
	"MedicCellularReactor": {
		Name: "Cellular Reactor",
		Text: "Consume Energy to heal",
	},
	"MedicClear": {
		Name: "Clear!",
		Text: "Increased area, bigger knockback",
	},
	"MedicEMPGrenade": {
		Name: "EMP Grenade",
		Text: "Displacement Grenade deals more damage to Shields",
	},
	"MedicExtendedCare": {
		Name: "Extended Care",
		Text: "Increase Healing Beam range",
	},
	"MedicFirstResponder": {
		Name: "First Responder",
		Text: "Healing Beam empowered at high Energy",
	},
	"MedicHeroicAbilityMedivacDropship": {
		Name: "Medivac Dropship",
		Text: "Transport allies across long distances",
	},
	"MedicHeroicAbilityStimDrone": {
		Name: "Stim Drone",
		Text: "Grant massive Attack and Movement Speed",
	},
	"MedicHyperactivity": {
		Name: "Hyperactivity",
		Text: "Lower cooldown, increase Movement Speed",
	},
	"MedicLifeSupport": {
		Name: "Life Support",
		Text: "Safeguard generates Energy",
	},
	"MedicPhysicalTherapy": {
		Name: "Physical Therapy",
		Text: "Safeguard removes Slows",
	},
	"MedicProlongedSafeguard": {
		Name: "Prolonged Safeguard",
		Text: "Increase Safeguard duration",
	},
	"MedicReinforcements": {
		Name: "Reinforcements",
		Text: "Medivac appears at Altar, reduce cooldown",
	},
	"MedicSafeZone": {
		Name: "Safe Zone",
		Text: "Exiting Medivac grants Protected",
	},
	"MedicSecondOpinion": {
		Name: "Second Opinion",
		Text: "Reduce Displacement Grenade cooldown",
	},
	"MedicShieldSequencer": {
		Name: "Shield Sequencer",
		Text: "Safeguard gains a second charge",
	},
	"MedicSystemShock": {
		Name: "System Shock",
		Text: "Displacement Grenade reduces damage dealt",
	},
	"MedicTraumaTrigger": {
		Name: "Trauma Trigger",
		Text: "Gain Armor at low Health",
	},
	"MedicVanadiumPlating": {
		Name: "Vanadium Plating",
		Text: "Stuns increase Safeguard Armor, duration",
	},
	"MedivhArcaneBrilliance": {
		Name: "Arcane Brilliance",
		Text: "Grant Spell Power and Mana to allies",
	},
	"MedivhArcaneRiftArcaneCharge": {
		Name: "Arcane Charge",
		Text: "Arcane Rift deals more damage after hitting a Hero",
	},
	"MedivhArcaneRiftGuardianOfTirisfal": {
		Name: "Guardian of Tirisfal",
		Text: "Arcane Rift kills Minions instantly",
	},
	"MedivhArcaneRiftManaAdept": {
		Name: "Mana Adept",
		Text: "Increases Arcane Rift Mana refund",
	},
	"MedivhArcaneRiftTheMastersTouch": {
		Name: "The Master's Touch",
		Text: "Quest: Hit  Heroes to increase damage",
	},
	"MedivhDustOfAppearance": {
		Name: "Dust of Appearance",
		Text: "Reveal Heroes in the surrounding area",
	},
	"MedivhForceOfWillArcaneExplosion": {
		Name: "Arcane Explosion",
		Text: "Force of Will deals area damage",
	},
	"MedivhForceOfWillCircleOfProtection": {
		Name: "Circle of Protection",
		Text: "Force of Will shields allies near the target",
	},
	"MedivhForceOfWillEnduringWill": {
		Name: "Enduring Will",
		Text: "Increases Force of Will duration",
	},
	"MedivhForceOfWillReabsorption": {
		Name: "Reabsorption",
		Text: "Force of Will heals the target",
	},
	"MedivhInvisibility": {
		Name: "Invisibility",
		Text: "Stealth an allied Hero",
	},
	"MedivhLeyLineSeal": {
		Name: "Ley Line Seal",
		Text: "Unleash a wave that Time Stops enemy Heroes",
	},
	"MedivhLeyLineSealMedivhCheats": {
		Name: "Medivh Cheats!",
		Text: "Ley Line Seal can be redirected while active",
	},
	"MedivhPolyBomb": {
		Name: "Poly Bomb",
		Text: "Polymorph a Hero that spreads to other Heroes",
	},
	"MedivhPolyBombGlyphOfPolyBomb": {
		Name: "Glyph Of Poly Bomb",
		Text: "Poly Bomb spreads faster and farther",
	},
	"MedivhPortalAstralProjection": {
		Name: "Astral Projection",
		Text: "Increases Portal range",
	},
	"MedivhPortalMageArmor": {
		Name: "Mage Armor",
		Text: "Using a Portal grants Physical Armor",
	},
	"MedivhPortalPortalMastery": {
		Name: "Portal Mastery",
		Text: "Place both Portal locations",
	},
	"MedivhPortalQuickening": {
		Name: "Quickening",
		Text: "Reduces Portal cooldown",
	},
	"MedivhPortalRavenFamiliar": {
		Name: "Raven Familiar",
		Text: "Using a Portal grants a Raven Familiar",
	},
	"MedivhPortalStablePortal": {
		Name: "Stable Portal",
		Text: "Increases Portal duration",
	},
	"MedivhTransformRavenBirdsEyeView": {
		Name: "Bird's Eye View",
		Text: "Increases Raven Form vision range",
	},
	"MedivhTransformRavenRavensIntellect": {
		Name: "Raven’s Intellect",
		Text: "Raven Form increases Mana Regen",
	},
	"MedivhTransformRavenWindsOfCelerity": {
		Name: "Winds of Celerity",
		Text: "Increases Raven Form Movement Speed",
	},
	"MonkAirAlly": {
		Name: "Air Ally",
		Text: "Place an Air Ally that reveals an area",
	},
	"MonkBlazingFistsDeadlyReach": {
		Name: "Blazing Fists",
		Text: "Basic Attacks reduce Deadly Reach cooldown",
	},
	"MonkBlindingSpeedRadiantDash": {
		Name: "Blinding Speed",
		Text: "Radiant Dash more often",
	},
	"MonkCleansingTouchRadiantDash": {
		Name: "Cleansing Touch",
		Text: "Radiant Dash makes ally Unstoppable",
	},
	"MonkDashofLight": {
		Name: "Dash of Light",
		Text: "Radiant Dash increases Breath of Heaven healing",
	},
	"MonkEarthAlly": {
		Name: "Earth Ally",
		Text: "Place an Earth Ally that grants Physical Armor",
	},
	"MonkEchoofHeavenBreathofHeaven": {
		Name: "Echo of Heaven",
		Text: "Breath of Heaven heals a second time",
	},
	"MonkElevenSidedStrikeSevenSidedStrike": {
		Name: "Transgression",
		Text: "Seven-Sided Strike hits more",
	},
	"MonkEpiphany": {
		Name: "Epiphany",
		Text: "Activate to restore Mana and Radiant Dash charges",
	},
	"MonkFistsofFuryDeadlyReach": {
		Name: "Fists of Fury",
		Text: "Increases Deadly Reach duration",
	},
	"MonkHeavenlyZealBreathofHeaven": {
		Name: "Heavenly Zeal",
		Text: "Breath of Heaven increases allies Move Speed",
	},
	"MonkHeroicAbilityDivinePalm": {
		Name: "Divine Palm",
		Text: "Prevent an ally from death",
	},
	"MonkHeroicAbilitySevenSidedStrike": {
		Name: "Seven-Sided Strike",
		Text: "Become Invulnerable and strike nearby Heroes",
	},
	"MonkInsight": {
		Name: "Insight",
		Text: "Quest: Basic Attacks restore mana, lower cooldowns",
	},
	"MonkIronFists": {
		Name: "Iron Fists",
		Text: "Basic Attacks increase Move Speed, damage",
	},
	"MonkPeacefulReposeDivinePalm": {
		Name: "Peaceful Repose",
		Text: "Reduces Divine Palm cooldown if they do not die",
	},
	"MonkQuicksilverRadiantDash": {
		Name: "Quicksilver",
		Text: "Radiant Dash gives Move Speed",
	},
	"MonkSixthSense": {
		Name: "Sixth Sense",
		Text: "Gain Physical Armor while Stunned, Rooted",
	},
	"MonkSpiritAlly": {
		Name: "Spirit Ally",
		Text: "Place a Spirit Ally that heals in an area",
	},
	"MonkTranscendence": {
		Name: "Transcendence",
		Text: "Basic Attacks heal",
	},
	"MonkWayoftheHundredFistsRadiantDash": {
		Name: "Way of the Hundred Fists",
		Text: "Radiant Dash launches a volley of blows",
	},
	"MuradinBronzebeardRage": {
		Name: "Bronzebeard Rage",
		Text: "Deals damage to nearby enemies",
	},
	"MuradinCombatStyleSkullcracker": {
		Name: "Skullcracker",
		Text: "Successive Basic Attacks Stun, deal more damage",
	},
	"MuradinCombatStyleThirdWind": {
		Name: "Third Wind",
		Text: "Increases Second Wind's healing",
	},
	"MuradinGiveEmTheAxeExecutioner60DamageBonus": {
		Name: "Give 'em the Axe!",
		Text: "Attacking disabled Heroes increases damage",
	},
	"MuradinHeroicAbilityAvatar": {
		Name: "Avatar",
		Text: "Temporary Health boost",
	},
	"MuradinHeroicAbilityHaymaker": {
		Name: "Haymaker",
		Text: "Stun and knock an enemy away",
	},
	"MuradinMasteryAvatarUnstoppableForce": {
		Name: "Unstoppable Force",
		Text: "Avatar grants additional bonuses",
	},
	"MuradinMasteryDwarfTossDwarfLaunch": {
		Name: "Dwarf Launch",
		Text: "Increase Dwarf Toss range",
	},
	"MuradinMasteryDwarfTossHeavyImpact": {
		Name: "Heavy Impact",
		Text: "Dwarf Toss briefly stuns enemies hit",
	},
	"MuradinMasteryDwarfTossLandingMomentum": {
		Name: "Landing Momentum",
		Text: "[PH] Increases the range and adds Movement Speed to Dwarf Toss",
	},
	"MuradinMasteryHaymakerGrandSlam": {
		Name: "Grand Slam",
		Text: "Increases Haymaker damage and adds charge",
	},
	"MuradinMasteryPassiveStoneform": {
		Name: "Stoneform",
		Text: "Activate to heal over time",
	},
	"MuradinMasteryStormhammerSledgehammer": {
		Name: "Sledgehammer",
		Text: "Increases Storm Bolt damage to non-Heroes",
	},
	"MuradinMasteryThunderburn": {
		Name: "Thunder Burn",
		Text: "Thunder Clap can strike a second time",
	},
	"MuradinMasteryThunderclapCrowdControl": {
		Name: "Crowd Control",
		Text: "Reduces Thunder Clap cooldown per target hit",
	},
	"MuradinMasteryThunderclapHealingStatic": {
		Name: "Healing Static",
		Text: "Thunder Clap heals Muradin",
	},
	"MuradinMasteryThunderclapReverberation": {
		Name: "Reverberation",
		Text: "Increases duration and Attack Speed reduction",
	},
	"MuradinMasteryThunderclapThunderstrike": {
		Name: "Thunder Strike",
		Text: "Increases Thunder Clap damage to one enemy",
	},
	"MuradinStormboltPerfectStorm": {
		Name: "Perfect Storm",
		Text: "Quest: Hitting Heroes empowers Storm Bolt",
	},
	"MuradinStormboltPiercingBolt": {
		Name: "Piercing Bolt",
		Text: "Storm Bolt pierces",
	},
	"MurkyAFishyDeal": {
		Name: "A Fishy Deal",
		Text: "Kill Minions with Pufferfish to bribe Mercs",
	},
	"MurkyBigTunaKahuna": {
		Name: "Big Tuna Kahuna",
		Text: "Increase Health and respawn time",
	},
	"MurkyBlackLagoon": {
		Name: "Black Lagoon",
		Text: "Increases Slime radius",
	},
	"MurkyEggHunt": {
		Name: "Egg Hunt",
		Text: "Place fake Eggs",
	},
	"MurkyEggShell": {
		Name: "Egg Shell",
		Text: "Spawning from Egg grants Shield",
	},
	"MurkyFishEye": {
		Name: "Fish Eye",
		Text: "Egg sees further, detects Stealth",
	},
	"MurkyFishOil": {
		Name: "Fish Oil",
		Text: "Casts Slime at Pufferfish location",
	},
	"MurkyFishTank": {
		Name: "Fish Tank",
		Text: "Basic Attacks grant Physical Armor",
	},
	"MurkyHeroicAbilityMarchoftheMurlocs": {
		Name: "March of the Murlocs",
		Text: "Swarm enemies with Murlocs",
	},
	"MurkyHeroicAbilityOctoGrab": {
		Name: "Octo-Grab",
		Text: "Stun and poke an enemy Hero",
	},
	"MurkyLivingtheDream": {
		Name: "Living the Dream",
		Text: "Quest: Living increases Spell Power",
	},
	"MurkyMakingInky": {
		Name: "Making Inky",
		Text: "Reduced Slime cooldown",
	},
	"MurkyMasteryAndASharkToo": {
		Name: "... And a Shark Too!",
		Text: "Increase Octo-Grab's damage",
	},
	"MurkyMasteryNeverEndingMurlocs": {
		Name: "Never-Ending Murlocs",
		Text: "March of the Murlocs can channel forever",
	},
	"MurkyMasteryRejuvenatingBubble": {
		Name: "Rejuvenating Bubble",
		Text: "Safety Bubble heals",
	},
	"MurkyMasteryTufferfish": {
		Name: "Tufferfish",
		Text: "Pufferfish gains Spell Armor, damage",
	},
	"MurkyMasteryWrathOfCod": {
		Name: "Wrath of Cod",
		Text: "Increase Pufferfish damage on Heroes",
	},
	"MurkySlimeTime": {
		Name: "Slime Time",
		Text: "Quest: Slime Heroes to empower Slime",
	},
	"MurkySlipperyWhenWet": {
		Name: "Slippery When Wet",
		Text: "Move faster during Safety Bubble, reduce cooldown",
	},
	"MurkyTimetoKrill": {
		Name: "Time to Krill",
		Text: "Basic Attacks slow enemies",
	},
	"MurkyToxicBuildup": {
		Name: "Toxic Buildup",
		Text: "Successive Basic Attacks cast Slime",
	},
	"NecromancerHeroicAbilityPoisonNova": {
		Name: "Poison Nova",
		Text: "Poison nearby enemies",
	},
	"NecromancerHeroicAbilitySkeletalMages": {
		Name: "Skeletal Mages",
		Text: "Summon Frost Mages to damage and Slow enemies",
	},
	"NecromancerTalentAmplifyDamage": {
		Name: "Amplify Damage",
		Text: "Bone Prison lowers Armor",
	},
	"NecromancerTalentAndarielsVisage": {
		Name: "Andariel's Visage",
		Text: "Poison Nova heals, increase Skeleton damage",
	},
	"NecromancerTalentBoneArmorBacklash": {
		Name: "Backlash",
		Text: "Bone Armor explodes for damage",
	},
	"NecromancerTalentBoneArmorShackler": {
		Name: "Shackler",
		Text: "Bone Armor Slows, can reduce cooldown",
	},
	"NecromancerTalentBoneArmorShade": {
		Name: "Shade",
		Text: "Bone Armor evades Basic Attacks",
	},
	"NecromancerTalentBoneSpear": {
		Name: "Bone Spear",
		Text: "Damage enemies in a line",
	},
	"NecromancerTalentColdHandOfDeath": {
		Name: "Cold Hand of Death",
		Text: "Increase Skeletal Mages Slow, Skeleton attacks Slow",
	},
	"NecromancerTalentCorpseExplosion": {
		Name: "Corpse Explosion",
		Text: "Skeletal Warriors explode on death",
	},
	"NecromancerTalentEchoesOfDeath": {
		Name: "Echoes of Death",
		Text: "Spectral Scythe spawns additional scythes",
	},
	"NecromancerTalentGrimScythe": {
		Name: "Grim Scythe",
		Text: "Reduce Cursed Strikes cooldown",
	},
	"NecromancerTalentHarvestVitality": {
		Name: "Harvest Vitality",
		Text: "Cursed Strikes against Heroes heals",
	},
	"NecromancerTalentJailors": {
		Name: "Jailors",
		Text: "Quest: Bone Prison spawns Skeletal Warriors",
	},
	"NecromancerTalentKalansEdict": {
		Name: "Kalan's Edict",
		Text: "Skeletal Warrior attacks reduce Heroic cooldown",
	},
	"NecromancerTalentMortalWound": {
		Name: "Mortal Wound",
		Text: "Spectral Scythe reduces healing taken",
	},
	"NecromancerTalentRapidHarvest": {
		Name: "Rapid Harvest",
		Text: "Cursed Strikes increases Attack Speed",
	},
	"NecromancerTalentReapersToll": {
		Name: "Reaper's Toll",
		Text: "Quest: Lower Spectral Scythe cooldown",
	},
	"NecromancerTalentTragOulsEssence": {
		Name: "Trag'Oul's Essence",
		Text: "Skeletal Warrior attacks restore Health, Mana",
	},
	"NecromancerTalentWeaken": {
		Name: "Weaken",
		Text: "Increase Cursed Strikes Attack Speed reduction",
	},
	"NovaAdvancedCloaking": {
		Name: "Advanced Cloaking",
		Text: "Stealth grants Mana, Move Speed ",
	},
	"NovaCombatStyleAntiArmorShells": {
		Name: "Anti-Armor Shells",
		Text: "Basic Attacks hit harder, reduce Physical Armor",
	},
	"NovaCombatStyleOneintheChamber": {
		Name: "One in the Chamber",
		Text: "Using an Ability empowers next Basic Attack",
	},
	"NovaCovertMission": {
		Name: "Covert Mission",
		Text: "Kill Minions or Heroes to bribe a Mercenary",
	},
	"NovaCovertOps": {
		Name: "Covert Ops",
		Text: "Stealth enhances Pinning Shot slow and cost",
	},
	"NovaGhostProtocol": {
		Name: "Ghost Protocol",
		Text: "Activate to gain Stealth",
	},
	"NovaHeroicAbilityPrecisionStrike": {
		Name: "Precision Strike",
		Text: "Area attack with unlimited range",
	},
	"NovaHeroicAbilityTripleTap": {
		Name: "Triple Tap",
		Text: "Fire at a target 3 times",
	},
	"NovaHoloStability": {
		Name: "Holo Stability",
		Text: "Increase Holo Decoy range, duration",
	},
	"NovaMasteryAmbushSnipe": {
		Name: "Ambush Snipe",
		Text: "Increased Snipe range while Cloaked",
	},
	"NovaMasteryCripplingShot": {
		Name: "Crippling Shot",
		Text: "Pinning Shot lowers Armor",
	},
	"NovaMasteryDigitalShrapnel": {
		Name: "Digital Shrapnel",
		Text: "Digital Shrapnel Reduces the cooldown of Holo Decoy by 2 seconds. Holo Decoy explodes on expiration, dealing  damage to nearby enemies.",
	},
	"NovaMasteryDoubleTap": {
		Name: "Double Tap",
		Text: "Pinning Shot now has 2 charges",
	},
	"NovaMasteryExplosiveShot": {
		Name: "Explosive Round",
		Text: "Snipe deals damage in an area",
	},
	"NovaMasteryFastReload": {
		Name: "Fast Reload",
		Text: "Triple Tap cooldown resets on Hero kills",
	},
	"NovaMasteryLethalDecoy": {
		Name: "Lethal Decoy",
		Text: "Holo Decoy does damage",
	},
	"NovaMasteryLongshot": {
		Name: "Longshot",
		Text: "Increases Pinning Shot range",
	},
	"NovaMasteryPerfectShotSnipe": {
		Name: "Perfect Shot",
		Text: "Reduces Snipe cooldown versus Heroes",
	},
	"NovaMasteryPrecisionBarrage": {
		Name: "Precision Barrage",
		Text: "Precision Strike can hold 2 charges",
	},
	"NovaMasteryPsiOpRangefinder": {
		Name: "Psi-Op Rangefinder",
		Text: "Increases Snipe's Range and lowers Cooldown",
	},
	"NovaMasteryPsionicEfficiency": {
		Name: "Psionic Efficiency",
		Text: "Removes Snipe Mana cost and increases range",
	},
	"NovaMasteryTazerRounds": {
		Name: "Tazer Rounds",
		Text: "Increased Pinning Shot slow duration",
	},
	"NovaNewRemoteDelivery": {
		Name: "Remote Delivery",
		Text: "Increases Holo Decoy cast and sight range",
	},
	"NovaRailgun": {
		Name: "Railgun",
		Text: "Snipe penetrates through enemies and cooldown is reduced",
	},
	"NovaRapidProjection": {
		Name: "Rapid Projection",
		Text: "Reduces Holo Decoy cooldown and Mana cost",
	},
	"NovaSnipeMaster": {
		Name: "Snipe Master",
		Text: "Increases Snipe damage from hitting Heroes",
	},
	"NovaTacticalEspionage": {
		Name: "Tactical Espionage",
		Text: "Increases Mana regen while Cloaked",
	},
	"PlaceholderRangedCombatStyle": {
		Name: "Ranged Placeholder",
		Text: "Placeholder Combat Style+20% Basic Attack+1 Range",
	},
	"ProbiusAggressiveMatrixWarpInPylon": {
		Name: "Aggressive Matrix",
		Text: "Power Fields give Attack Damage",
	},
	"ProbiusConstructAdditionalPylonsPylonOvercharge": {
		Name: "Construct Additional Pylons",
		Text: "Increase max Pylon count and Pylon Overcharge damage",
	},
	"ProbiusEchoPulseDisruptionPulse": {
		Name: "Echo Pulse",
		Text: "Disruption Pulse returns to Probius",
	},
	"ProbiusGateKeeperNullGate": {
		Name: "Gate Keeper",
		Text: "Indefinite duration while in Pylon Power",
	},
	"ProbiusGatherMineralsPhotonCannonQuest": {
		Name: "Gather Minerals",
		Text: "Quest: Collect Minerals to empower Photon Cannons",
	},
	"ProbiusGravityWellWarpRift": {
		Name: "Gravity Well",
		Text: "Increase Warp Rift Slow amount in center",
	},
	"ProbiusInterferenceWarpRift": {
		Name: "Interference",
		Text: "Warp Rift explosions reduce enemy Spell Power",
	},
	"ProbiusNullGateHeroic": {
		Name: "Null Gate",
		Text: "Create a slowing barrier",
	},
	"ProbiusParticleAcceleratorDisruptionPulse": {
		Name: "Particle Accelerator",
		Text: "Increase damage for each target hit",
	},
	"ProbiusPhotonBarrierPhotonCannon": {
		Name: "Photon Barrier",
		Text: "Gain Spell Armor while Photon Cannon is Active",
	},
	"ProbiusPowerOverflowingWarpInPylon": {
		Name: "Power Overflowing",
		Text: "Power Fields give Spell Power, Mana",
	},
	"ProbiusProbiusLoopWarpRift": {
		Name: "Probius Loop",
		Text: "Create new Warp Rifts by hitting Heroes with explosions",
	},
	"ProbiusPylonOverchargeHeroic": {
		Name: "Pylon Overcharge",
		Text: "Increase the size of Pylon power fields and allow them to attack enemies within it for  damage per second. Lasts  seconds.",
	},
	"ProbiusQuantumEntanglementWarpRift": {
		Name: "Quantum Entanglement",
		Text: "Warp Rift's slow lingers",
	},
	"ProbiusRepulsorWarpRift": {
		Name: "Repulsor",
		Text: "Enemies are knocked away from Warp Rift detonations",
	},
	"ProbiusRiftShockWarpRift": {
		Name: "Rift Shock",
		Text: "Deal more damage to enemies caught in Warp Rift detonations",
	},
	"ProbiusShieldBatteryWarpInPylon": {
		Name: "Shield Battery",
		Text: "Power Fields grant Shields",
	},
	"ProbiusShieldCapacitor": {
		Name: "Shield Capacitor",
		Text: "Gain permanent Shields",
	},
	"ProbiusShootEmUptoRiftsDisruptionPulse": {
		Name: "Shoot 'Em Up",
		Text: "Hitting a Warp Rift splits Disruption Pulse into several additional pulses",
	},
	"ProbiusTowerDefensePhotonCannon": {
		Name: "Tower Defense",
		Text: "Reduce Photon Cannon cooldown",
	},
	"ProbiusTurboChargedWorkerRush": {
		Name: "Turbo Charged",
		Text: "Move faster and decrease Worker Rush's cooldown",
	},
	"ProbiusWarpResonanceWarpRiftQuest": {
		Name: "Warp Resonance",
		Text: " Quest: Detonate Warp Rift on Heroes",
	},
	"RagnarosBlastWaveBlastEcho": {
		Name: "Blast Echo",
		Text: "A second Blast Wave is created after the first",
	},
	"RagnarosBlastWaveEngulfingFlame": {
		Name: "Engulfing Flame",
		Text: "Increase Blast Wave damage, radius",
	},
	"RagnarosBlastWaveSlowBurn": {
		Name: "Slow Burn",
		Text: "Blast Wave slows enemies",
	},
	"RagnarosBlastWaveSuperheated": {
		Name: "Superheated",
		Text: "Blast Wave damage increased when used on Ragnaros",
	},
	"RagnarosBlastWaveTemperedFlame": {
		Name: "Tempered Flame",
		Text: "Blast Wave damage gives Shields",
	},
	"RagnarosBlisteringAttacks": {
		Name: "Blistering Attacks",
		Text: "Periodically empower Basic Attacks against Heroes",
	},
	"RagnarosCatchingFire": {
		Name: "Catching Fire",
		Text: "Quest: Regen Globes increase Health Regen",
	},
	"RagnarosEmpowerSulfurasCauterizeWounds": {
		Name: "Cauterize Wounds",
		Text: "Empower Sulfuras healing increased",
	},
	"RagnarosEmpowerSulfurasGiantScorcher": {
		Name: "Giant Scorcher",
		Text: "Empower Sulfuras deals extra damage to Heroes",
	},
	"RagnarosEmpowerSulfurasHandOfRagnaros": {
		Name: "Hand of Ragnaros",
		Text: "Reduce Empower Sulfuras cooldown",
	},
	"RagnarosEmpowerSulfurasSulfurasHungers": {
		Name: "Sulfuras Hungers",
		Text: "Quest: Kill Minions with Empower Sulfuras to increase its damage",
	},
	"RagnarosLavaWave": {
		Name: "Lava Wave",
		Text: "Send a wave of lava down a lane",
	},
	"RagnarosLavaWaveLavaSurge": {
		Name: "Lava Surge",
		Text: "Lava Wave has two charges",
	},
	"RagnarosLivingMeteorFireWard": {
		Name: "Fire Ward",
		Text: "Living Meteor grants Spell Block",
	},
	"RagnarosLivingMeteorMeteorBomb": {
		Name: "Meteor Bomb",
		Text: "Living Meteor explodes at the end of its path",
	},
	"RagnarosLivingMeteorMoltenPower": {
		Name: "Molten Power",
		Text: "After hitting Heroes, next Living Meteor deals increased damage",
	},
	"RagnarosLivingMeteorShiftingMeteor": {
		Name: "Shifting Meteor",
		Text: "Quest: Hit Heroes to empower Living Meteor",
	},
	"RagnarosMoltenCoreHeroicDifficulty": {
		Name: "Heroic Difficulty",
		Text: "Enhances Molten Core Health, cooldown, damage",
	},
	"RagnarosResilientFlame": {
		Name: "Resilient Flame",
		Text: "Gain Armor when Stunned",
	},
	"RagnarosSubmerge": {
		Name: "Submerge",
		Text: "Activate to heal and enter Stasis",
	},
	"RagnarosSulfurasSmash": {
		Name: "Sulfuras Smash",
		Text: "Damage and Stun enemies in an area",
	},
	"RagnarosSulfurasSmashFlamesOfSulfuron": {
		Name: "Flames of Sulfuron",
		Text: "Sulfuras Smash slows enemies, stuns longer",
	},
	"RaynorACardToPlay": {
		Name: "A Card to Play",
		Text: "Hero deaths lower Heroic cooldown",
	},
	"RaynorHeroicAbilityHyperion": {
		Name: "Hyperion",
		Text: "Call down a strafing Battlecruiser",
	},
	"RaynorHeroicAbilityRaynorsRaiders": {
		Name: "Raynor's Raiders",
		Text: "Summon a pair of Banshee escorts",
	},
	"RaynorHyperionGroundStrafe": {
		Name: "Scorched Earth",
		Text: "More shots dealing area damage",
	},
	"RaynorInspireConfidentAim": {
		Name: "Confident Aim",
		Text: "Reduces Penetrating Round cooldown",
	},
	"RaynorInspireSteelResolve": {
		Name: "Steel Resolve",
		Text: "Inspire lasts longer",
	},
	"RaynorLeadFromTheFrontPuttinOnAClinic": {
		Name: "Puttin' On a Clinic",
		Text: "Reduces cooldowns when killing enemy minions",
	},
	"RaynorMasteryActivatedRushAdrenalineRush": {
		Name: "Activated Rush",
		Text: "Lowers the cooldown of Adrenaline Rush by  seconds, and it can be manually activated.",
	},
	"RaynorMasteryBattleHyperionHyperion": {
		Name: "Battle Hyperion",
		Text: "Battle Hyperion Number of shots on each target increased to 2 per volley. Every  seconds Hyperion will fire its Yamato cannon at enemy structures.",
	},
	"RaynorMasteryBullseyePenetratingRound": {
		Name: "Bullseye",
		Text: "Penetrating Round stuns first enemy hit",
	},
	"RaynorMasteryClusterRoundPenetratingRound": {
		Name: "Cluster Round",
		Text: "Increase Penetrating Round damage, width",
	},
	"RaynorMasteryFightorFlightAdrenalineRush": {
		Name: "Fight or Flight",
		Text: "Adrenaline Rush activateable, grants Armor",
	},
	"RaynorMasteryGiveMeMoreAdrenalineRush": {
		Name: "Give Me More!",
		Text: "Adrenaline Rush heals more",
	},
	"RaynorMasteryHelsAngelsRaynorsBanshees": {
		Name: "Dusk Wings",
		Text: "Banshees stay cloaked, attack faster",
	},
	"RaynorMasteryInspireRevolutionOverdrive": {
		Name: "Revolution Overdrive",
		Text: "Gain Movement Speed per inspired Hero",
	},
	"RaynorPenetratingRoundDoubleBarreled": {
		Name: "Double-Barreled",
		Text: "Penetrating Round gains a second charge",
	},
	"RaynorPenetratingRoundHamstringShot": {
		Name: "Hamstring Shot",
		Text: "Penetrating Round slows",
	},
	"RaynorRelentlessLeader": {
		Name: "Relentless Leader",
		Text: "Knock enemies back when Stunned or Rooted",
	},
	"RehgarEarthlivingEnchant": {
		Name: "Earthliving Enchant",
		Text: "Empowers Chain Heal on low Health Heroes",
	},
	"RehgarGhostWolfBloodAndThunder": {
		Name: "Blood and Thunder",
		Text: "Ghost Wolf attacks reduce Basic Ability cooldowns",
	},
	"RehgarHeroicAbilityAncestralHealing": {
		Name: "Ancestral Healing",
		Text: "Significantly heal allied Hero",
	},
	"RehgarHeroicAbilityBloodlust": {
		Name: "Bloodlust",
		Text: "Give allied Heroes Movement Speed, Attack Speed, and Life Steal",
	},
	"RehgarHungeroftheWolf": {
		Name: "Hunger of the Wolf",
		Text: "Ghost Wolf attacks against Heroes deal an additional % of the target's maximum Health and heal Rehgar for % of his maximum Health.",
	},
	"RehgarMasteryChainReaction": {
		Name: "Chain Reaction",
		Text: "Increases Chain Heal on Lightning Shielded allies",
	},
	"RehgarMasteryColossalTotem": {
		Name: "Colossal Totem",
		Text: "Increases Earthbind Totem area and range",
	},
	"RehgarMasteryDeepHealing": {
		Name: "Deep Healing",
		Text: "Deep Healing Increases Chain Heal's healing by 30% on targets that are below 50% health.",
	},
	"RehgarMasteryEarthGraspTotem": {
		Name: "Earthgrasp Totem",
		Text: "Increases Earthbind Totem initial slow",
	},
	"RehgarMasteryEarthShield": {
		Name: "Earth Shield",
		Text: "Lightning Shield gives a Shield",
	},
	"RehgarMasteryElectricCharge": {
		Name: "Electric Charge",
		Text: "Increases Lightning Shield radius",
	},
	"RehgarMasteryFarseersBlessing": {
		Name: "Farseer's Blessing",
		Text: "Increases heal and heals in area",
	},
	"RehgarMasteryFarsight": {
		Name: "Farsight",
		Text: "Activate to reveal target area and nearby enemies",
	},
	"RehgarMasteryFeralHeart": {
		Name: "Feral Heart",
		Text: "Increases Health and Mana Regen in Ghost Wolf",
	},
	"RehgarMasteryGladiatorsWarShout": {
		Name: "Gladiator's War Shout",
		Text: "Increases Bloodlust Life Steal",
	},
	"RehgarMasteryHealingSurge": {
		Name: "Healing Surge",
		Text: "Increases Chain Heal amount and heal an additional target",
	},
	"RehgarMasteryLightningBond": {
		Name: "Lightning Bond",
		Text: "Lightning Shield is also cast on Rehgar",
	},
	"RehgarMasteryLightningTotem": {
		Name: "Lightning Totem",
		Text: "Lightning Totem Earthbind Totem automatically gains Lightning Shield when created.",
	},
	"RehgarMasteryShamanHealingWard": {
		Name: "Healing Totem",
		Text: "Activate to place a ward which heals in an area",
	},
	"RehgarMasterySpiritwalkersGrace": {
		Name: "Spiritwalker's Grace",
		Text: "Reduces Chain Heal Mana cost",
	},
	"RehgarMasteryStormcaller": {
		Name: "Stormcaller",
		Text: "Lightning Shield returns Mana",
	},
	"RehgarMasteryTidalWaves": {
		Name: "Tidal Waves",
		Text: "Reduces Chain Heal cooldown",
	},
	"RehgarMasteryWolfRun": {
		Name: "Wolf Run",
		Text: "Increases Ghost Wolf Move Speed",
	},
	"RehgarRisingStorm": {
		Name: "Rising Storm",
		Text: "Increases Lightning Shield damage with each hit",
	},
	"RehgarTotemicProjection": {
		Name: "Totemic Projection",
		Text: "Earthbind Totem can be moved",
	},
	"RexxarAnimalHusbandry": {
		Name: "Animal Husbandry",
		Text: "Max Health increase over time",
	},
	"RexxarAspectOfTheHawkSpiritSwoop": {
		Name: "Aspect of the Hawk",
		Text: "Spirit Swoop increases Attack Speed",
	},
	"RexxarAspectoftheBeastCharge": {
		Name: "Aspect of the Beast",
		Text: "Misha's Basic Attacks lower Misha, Charge! cooldown",
	},
	"RexxarBarkskin": {
		Name: "Barkskin",
		Text: "Misha gains Spell Armor after charging",
	},
	"RexxarBearNecessitiesCharge": {
		Name: "Bear Necessities",
		Text: "Misha's Basic Attacks heal Rexxar",
	},
	"RexxarCriticalCareMendPet": {
		Name: "Critical Care",
		Text: "Increases Mend Pet healing when Misha is low",
	},
	"RexxarDireBeast": {
		Name: "Dire Beast",
		Text: "Basic Attacks increase damage of Misha, Charge!",
	},
	"RexxarFeignDeath": {
		Name: "Feign Death",
		Text: "Become Invulnerable and control Misha",
	},
	"RexxarFlare": {
		Name: "Flare",
		Text: "Gain vision of an area",
	},
	"RexxarFrenzyofKalimdor": {
		Name: "Frenzy of Kalimdor",
		Text: "Rexxar's Basic Attacks deal more damage, Misha's attacks slow",
	},
	"RexxarHardenedSkin": {
		Name: "Hardened Skin",
		Text: "Rexxar and Misha gain massive Armor",
	},
	"RexxarHeroicAbilityBestialWrath": {
		Name: "Bestial Wrath",
		Text: "Increases Misha's Basic Attacks",
	},
	"RexxarHeroicAbilityUnleashTheBoars": {
		Name: "Unleash the Boars",
		Text: "Send a pack of boars to hunt enemy Heroes",
	},
	"RexxarHungryBear": {
		Name: "Hungry Bear",
		Text: "Misha's Basic Attacks heal her",
	},
	"RexxarHunterGatherer": {
		Name: "Hunter-Gatherer",
		Text: "Quest: Regen Globes grant Health Regen",
	},
	"RexxarMendPetBloodOfTheRhino": {
		Name: "Blood of the Rhino",
		Text: "Increases Mend Pet duration",
	},
	"RexxarSpiritBond": {
		Name: "Spirit Bond",
		Text: "Misha's Basic Attacks heal Rexxar",
	},
	"RexxarSpiritBondEasyPrey": {
		Name: "Easy Prey",
		Text: "Misha is more effective against Minions and Mercs",
	},
	"RexxarSpiritBondGrizzledFortitude": {
		Name: "Grizzled Fortitude",
		Text: "Periodically gain Physical Armor",
	},
	"RexxarSpiritBondPrimalIntimidation": {
		Name: "Primal Intimidation",
		Text: "Activate to slow enemy Attack and Move speed",
	},
	"RexxarSpiritBondWildfireBear": {
		Name: "Wildfire Bear",
		Text: "Misha deals damage to nearby enemies",
	},
	"RexxarSpiritSwoopBirdOfPrey": {
		Name: "Bird of Prey",
		Text: "Increases Spirit Swoop damage to Minions",
	},
	"RexxarSpiritSwoopCripplingTalons": {
		Name: "Crippling Talons",
		Text: "Increases Spirit Swoop slow amount and duration",
	},
	"RexxarSurvivalistTraining": {
		Name: "Survivalist Training",
		Text: "Regen Globes restore more Mana",
	},
	"RexxarTakingFlightSpiritSwoop": {
		Name: "Taking Flight",
		Text: "Increases Spirit Swoop range, returns Mana",
	},
	"RexxarThrilloftheHunt": {
		Name: "Thrill of the Hunt",
		Text: "Basic Attacks increase Move Speed",
	},
	"RexxarUnleashTheBoarsKillCommand": {
		Name: "Kill Command",
		Text: "Unleash the Boars deals more damage and roots",
	},
	"SamuroAdvancingStrikesDeflection": {
		Name: "Deflection",
		Text: "Advancing Strikes grants Physical Armor",
	},
	"SamuroBlademastersPursuit": {
		Name: "Blademaster’s Pursuit",
		Text: "Increase Advancing Strikes effect and duration",
	},
	"SamuroBurningBlade": {
		Name: "Burning Blade",
		Text: "Critical Strikes damage nearby enemies",
	},
	"SamuroHarshWinds": {
		Name: "Harsh Winds",
		Text: "Attack from Stealth to increase damage",
	},
	"SamuroHeroicAbilityBladestorm": {
		Name: "Bladestorm",
		Text: "Become an unstoppable whirlwind of death",
	},
	"SamuroHeroicAbilityIllusionMaster": {
		Name: "Illusion Master",
		Text: "Images can be controlled",
	},
	"SamuroMercilessStrikes": {
		Name: "Merciless Strikes",
		Text: "Always Critically Strike disabled targets",
	},
	"SamuroMirage": {
		Name: "Mirage",
		Text: "Mirror Image grants Spell Armor",
	},
	"SamuroMirrorImageWayOfTheBlade": {
		Name: "Way of the Blade",
		Text: "Critical Strike happens more often",
	},
	"SamuroMirroredSteel": {
		Name: "Mirrored Steel",
		Text: "Attacks reduce Mirror Image cooldown",
	},
	"SamuroOverpoweringBlow": {
		Name: "Crushing Blows",
		Text: "Critical Strike has two charges",
	},
	"SamuroPhantomPain": {
		Name: "Phantom Pain",
		Text: "Images increase Critical Strike damage",
	},
	"SamuroPressTheAttack": {
		Name: "Press the Attack",
		Text: "Advancing Strikes increases Attack Speed",
	},
	"SamuroShukuchi": {
		Name: "Shukuchi",
		Text: "Wind Walk teleports a short distance",
	},
	"SamuroThreeBladeStyle": {
		Name: "Three Blade Style",
		Text: "Images have increased Health and Duration",
	},
	"SamuroWayOfIllusion": {
		Name: "Way of Illusion",
		Text: "Quest: Image Critical Strikes empower Basic Attacks",
	},
	"SamuroWayOfTheWind": {
		Name: "Way of the Wind",
		Text: " Quest: Attacks from Stealth empower Wind Walk",
	},
	"SamuroWhirlwindStorm": {
		Name: "Dance Of Death",
		Text: "Images use Bladestorm",
	},
	"SamuroWindStrider": {
		Name: "Wind Strider",
		Text: "Greatly reduce Wind Walk cooldown",
	},
	"SamuroWindwalkKawarimi": {
		Name: "Kawarimi",
		Text: "Wind Walk creates an Image",
	},
	"SamuroWindwalkOneWithTheWind": {
		Name: "One With The Wind",
		Text: "Gain Armor while Stealthed",
	},
	"SgtHammerAmbush": {
		Name: "Ambush",
		Text: "Stealth and extra damage in Siege Mode",
	},
	"SgtHammerBullheadMines": {
		Name: "Bullhead Mines",
		Text: "Middle Spider Mine knocks back",
	},
	"SgtHammerFirstStrike": {
		Name: "First Strike",
		Text: "Basic Attacks do extra damage if not attacked",
	},
	"SgtHammerGraduatingRange": {
		Name: "Graduating Range",
		Text: "Gradually increase range in Siege Mode",
	},
	"SgtHammerHeroicAbilityBluntForceGun": {
		Name: "Blunt Force Gun",
		Text: "Fire a large bullet across the map",
	},
	"SgtHammerHeroicAbilityNapalmStrike": {
		Name: "Napalm Strike",
		Text: "Fire damaging napalm at enemies",
	},
	"SgtHammerHyperCoolingEngines": {
		Name: "Hyper-Cooling Engines",
		Text: "Thrusters cooldown reduced",
	},
	"SgtHammerMasteryAdvancedArtillery": {
		Name: "Advanced Artillery",
		Text: "Increase damage to distant enemies",
	},
	"SgtHammerMasteryAdvancedLavaStrikeNapalmStrike": {
		Name: "Advanced Lava Strike",
		Text: "Increased range and impact damage",
	},
	"SgtHammerMasteryExcessiveForceConcussiveBlast": {
		Name: "Excessive Force",
		Text: "Double the knock back distance",
	},
	"SgtHammerMasteryFlakCannons": {
		Name: "Barricade",
		Text: "Concussive Blast creates a wall",
	},
	"SgtHammerMasteryHoverSiegeMode": {
		Name: "Hover Siege Mode",
		Text: "Move slowly in Siege Mode",
	},
	"SgtHammerMasteryLethalBlastConcussiveBlast": {
		Name: "Lethal Blast",
		Text: "Increase Concussive Blast damage",
	},
	"SgtHammerMasteryMaelstromShells": {
		Name: "Maelstrom Shells",
		Text: "Increase Basic Attack range",
	},
	"SgtHammerMasteryMineField": {
		Name: "Mine Field",
		Text: "Spawn 2 more mines.",
	},
	"SgtHammerMasteryOrbitalBFGBluntForceGun": {
		Name: "Orbital BFG",
		Text: "BFG repeatedly travels through the map",
	},
	"SgtHammerMasterySlowingMinesSpiderMines": {
		Name: "Slowing Mines",
		Text: "Increase slow amount and duration",
	},
	"SgtHammerResistant": {
		Name: "Resistant",
		Text: "Stuns, Roots give Armor in Siege Mode",
	},
	"SonyaTalentIgnorePain": {
		Name: "Ignore Pain",
		Text: "Activate to gain massive Armor",
	},
	"SonyaTalentNervesOfSteel": {
		Name: "Nerves of Steel",
		Text: "Activate to gain a Shield",
	},
	"StitchesCannibalize": {
		Name: "Cannibalize",
		Text: "Basic Attacks against Heroes heal",
	},
	"StitchesCombatStyleTenderizer": {
		Name: "Tenderizer",
		Text: "Basic Attacks slow enemies",
	},
	"StitchesCombatStyleVileCleaver": {
		Name: "Vile Cleaver",
		Text: "Basic Attacks create Vile Gas",
	},
	"StitchesFleaBag": {
		Name: "Flea Bag",
		Text: "Stuns and Roots reduce Cooldowns",
	},
	"StitchesHeroicAbilityGorge": {
		Name: "Gorge",
		Text: "Completely consume an enemy Hero",
	},
	"StitchesHeroicAbilityPutridBile": {
		Name: "Putrid Bile",
		Text: "Emit a slowing bile puddle",
	},
	"StitchesHungryforMore": {
		Name: "Hungry for More",
		Text: "Quest: Regen Globes increase Health",
	},
	"StitchesMasteryChewYourFood": {
		Name: "Chew Your Food",
		Text: "Devour heals for more over time",
	},
	"StitchesMasteryFishingHook": {
		Name: "Fishing Hook",
		Text: "Increased range",
	},
	"StitchesMasteryHeavySlam": {
		Name: "Heavy Slam",
		Text: "Increased Slam damage",
	},
	"StitchesMasteryHungryHungryStitchesGorge": {
		Name: "Hungry Hungry Stitches",
		Text: "Gorge multiple Heroes, can teleport",
	},
	"StitchesMasteryIndigestionDevour": {
		Name: "Indigestion",
		Text: "Devour creates a Retchling",
	},
	"StitchesMasteryLastBiteDevour": {
		Name: "Last Bite",
		Text: "Reduced cooldown on kill",
	},
	"StitchesMasteryMegaSmashSlam": {
		Name: "Mega Smash",
		Text: "Increases size of Slam",
	},
	"StitchesMasteryPulverizeSlam": {
		Name: "Pulverize",
		Text: "Reduces Slam cooldown and adds slow",
	},
	"StitchesMasteryPutridGroundSlam": {
		Name: "Putrid Ground",
		Text: "Slam infects with Vile Gas",
	},
	"StitchesMasterySavorTheFlavorDevour": {
		Name: "Savor the Flavor",
		Text: "Quest: Devouring Heroes gives permanent Health regen",
	},
	"StitchesMasteryShishKabobHook": {
		Name: "Shish Kabob",
		Text: "Hook pulls 2 targets",
	},
	"StitchesPotentBile": {
		Name: "Potent Bile",
		Text: "Putrid Bile lasts longer and slows more",
	},
	"StitchesRestorativeFumes": {
		Name: "Restorative Fumes",
		Text: "Vile Gas heals Stitches",
	},
	"StitchesToxicGas": {
		Name: "Toxic Gas",
		Text: "Increases area and damage of Vile Gas",
	},
	"StukovBioExplosionSwitch": {
		Name: "Bio-Explosion Switch",
		Text: "Bio Kill-Switch detonates Lurking Arm",
	},
	"StukovBioticArmor": {
		Name: "Biotic Armor",
		Text: "Healing Pathogen grants Physical Armor",
	},
	"StukovControlledChaos": {
		Name: "Controlled Chaos",
		Text: "Flailing Swipe only hits once, gains multiple charges",
	},
	"StukovEyeInfection": {
		Name: "Eye Infection",
		Text: "Increase Weighted Pustule damage, detonation Blinds",
	},
	"StukovFetidTouch": {
		Name: "Fetid Touch",
		Text: "Quest: Reduce Weighted Pustule cooldown, cost",
	},
	"StukovGrowingInfestation": {
		Name: "Growing Infestation",
		Text: "Lurking Arm expands, cooldown increased",
	},
	"StukovHeroicAbilityFlailingSwipe": {
		Name: "Flailing Swipe",
		Text: "Knock back enemies in front of Stukov",
	},
	"StukovHeroicAbilityMassiveShove": {
		Name: "Massive Shove",
		Text: "Shove an enemy until they collide with terrain",
	},
	"StukovItHungers": {
		Name: "It Hungers",
		Text: "Hitting Heroes reduces Lurking Arm cooldown",
	},
	"StukovLingeringSpines": {
		Name: "Lingering Spines",
		Text: "Lurking Arm persists after canceling",
	},
	"StukovLongPitch": {
		Name: "The Long Pitch",
		Text: "Increase Weighted Pustule range",
	},
	"StukovLowBlow": {
		Name: "Low Blow",
		Text: "Increase Lurking Arm damage to low Health targets",
	},
	"StukovOneGoodSpread": {
		Name: "One Good Spread...",
		Text: "Reduce Healing Pathogen cooldown after spreading",
	},
	"StukovPoppinPustules": {
		Name: "Poppin' Pustules",
		Text: "Quest: Empower Weighted Pustule",
	},
	"StukovPoxPopuli": {
		Name: "Pox Populi",
		Text: "Bio-Kill Switch no longer removes Healing Pathogen",
	},
	"StukovPushComesToShove": {
		Name: "Push Comes To Shove",
		Text: "Increase Massive Shove speed, can reduce cooldown",
	},
	"StukovSpineLauncher": {
		Name: "Spine Launcher",
		Text: "Basic Attacks Slow, become ranged",
	},
	"StukovSuperstrain": {
		Name: "Superstrain",
		Text: "Healing Pathogen heals when Stunned, Rooted",
	},
	"StukovTargetedExcision": {
		Name: "Targeted Excision",
		Text: "Detonate single Weighted Pustule to reduce cooldown",
	},
	"StukovTopOff": {
		Name: "Top Off",
		Text: "Healing Pathogen stronger on high Health target",
	},
	"StukovUniversalCarrier": {
		Name: "Universal Carrier",
		Text: "Healing Pathogen can spread continually",
	},
	"StukovVigorousReuptake": {
		Name: "Vigorous Reuptake",
		Text: "Quest: Increase Bio-Kill Switch healing",
	},
	"StukovVirulentReaction": {
		Name: "Virulent Reaction",
		Text: "Detonating Weighted Pustule in Lurking Arm Roots",
	},
	"StukovWeightedPustuleReactiveBallistospores": {
		Name: "Reactive Ballistospores",
		Text: "Radiate Weighted Pustule at low Health",
	},
	"StukovWithinMyReach": {
		Name: "Within My Reach",
		Text: "Increase Lurking Arm range",
	},
	"SylvanasDreadfulWake": {
		Name: "Dreadful Wake",
		Text: "Haunting Wave stuns Minions and Mercs longer",
	},
	"SylvanasHeroicAbilityMindControl": {
		Name: "Mind Control",
		Text: "Briefly control an enemy Hero",
	},
	"SylvanasHeroicAbilityPossession": {
		Name: "Possession",
		Text: "Convert enemy Minions",
	},
	"SylvanasHeroicAbilityWailingArrow": {
		Name: "Wailing Arrow",
		Text: "Fire arrow that damages and silences",
	},
	"SylvanasTalentBlackArrowsParalysis": {
		Name: "Paralysis",
		Text: "Increases Black Arrows duration",
	},
	"SylvanasTalentColdEmbrace": {
		Name: "Cold Embrace",
		Text: "Shadow Dagger lowers Armor",
	},
	"SylvanasTalentCorruption": {
		Name: "Corruption",
		Text: "Basic Attacks destroy Ammo",
	},
	"SylvanasTalentDarkLadysCallMindControl": {
		Name: "Dark Lady's Call",
		Text: "Increase Mind Controlled enemies Move Speed",
	},
	"SylvanasTalentLifeDrain": {
		Name: "Life Drain",
		Text: "Shadow Dagger heals",
	},
	"SylvanasTalentLostSoul": {
		Name: "Lost Soul",
		Text: "Reduces Shadow Dagger cooldown on Heroes",
	},
	"SylvanasTalentMercenaryQueen": {
		Name: "Mercenary Queen",
		Text: "Nearby Mercs are stronger",
	},
	"SylvanasTalentOverflowingQuiver": {
		Name: "Overflowing Quiver",
		Text: "Shoot a free Withering Fire on Minion kills",
	},
	"SylvanasTalentOverwhelmingAffliction": {
		Name: "Overwhelming Affliction",
		Text: "Black Arrows slows Heroes",
	},
	"SylvanasTalentRemorseless": {
		Name: "Remorseless",
		Text: "Abilities increase Basic Attack damage",
	},
	"SylvanasTalentUnstablePoison": {
		Name: "Unstable Poison",
		Text: "Black Arrows explodes Minions on death",
	},
	"SylvanasTalentWillOfTheForsaken": {
		Name: "Will of the Forsaken",
		Text: "Activate to gain Unstoppable and Move Speed",
	},
	"SylvanasTalentWindrunnerHauntingWave": {
		Name: "Windrunner",
		Text: "Haunting Wave recharges Withering Fire and is usable twice",
	},
	"SylvanasTalentWithTheWind": {
		Name: "With the Wind",
		Text: "Increases Withering Fire range",
	},
	"SylvanasTalentWitheringBarrage": {
		Name: "Withering Barrage",
		Text: "Withering Fire gains an additional charge and fires faster",
	},
	"SylvanasTalentWitheringFireBarbedShot": {
		Name: "Barbed Shot",
		Text: "Withering Fire deals triple damage to non-Heroic units.",
	},
	"SylvanasTalentWitheringFireEvasiveFire": {
		Name: "Evasive Fire",
		Text: "Withering Fire grants Move Speed",
	},
	"SylvanasWailingArrowDeafeningBlast": {
		Name: "Deafening Blast",
		Text: "Increases power of Wailing Arrow center",
	},
	"TalentNullificationShield": {
		Name: "Nullification Shield",
		Text: "Activate to gain massive Spell Armor",
	},
	"TassadarAdunsWisdom": {
		Name: "Adun's Wisdom",
		Text: "Oracle reduces Basic Ability cooldowns",
	},
	"TassadarDeepShift": {
		Name: "Deep Shift",
		Text: "Increase Dimensional Shift duration, speed",
	},
	"TassadarFocusedBeam": {
		Name: "Focused Beam",
		Text: "Oracle increases Basic Attack damage",
	},
	"TassadarHeroicAbilityArchon": {
		Name: "Phase Shift",
		Text: "Become an Archon",
	},
	"TassadarHeroicAbilityForceWall": {
		Name: "Force Wall",
		Text: "Create an unpathable wall",
	},
	"TassadarKhalasCelerityPlasmaShield": {
		Name: "Khala's Celerity",
		Text: "Plasma Shield gives Move Speed",
	},
	"TassadarKhalasEmbrace": {
		Name: "Khala's Embrace",
		Text: "Increased Plasma Shield Life Steal",
	},
	"TassadarKhalasLight": {
		Name: "Khala's Light",
		Text: "Plasma Shield grants Armor on expiration",
	},
	"TassadarKhaydarinResonance": {
		Name: "Khaydarin Resonance",
		Text: "Quest: Gather Regen Globes to empower Plasma Shield",
	},
	"TassadarMasteryForceBarrier": {
		Name: "Force Barrier",
		Text: "Force Wall's range, cooldown enhanced",
	},
	"TassadarMasteryMentalAcuity": {
		Name: "Mental Acuity",
		Text: "Quest: Takedowns reduce the Cooldown of Oracle",
	},
	"TassadarMasteryTwilightArchon": {
		Name: "Phase Shift",
		Text: "Basic Attacks increase Archon duration",
	},
	"TassadarNullification": {
		Name: "Nullification",
		Text: "Greatly lower an enemy Hero's damage",
	},
	"TassadarPhaseDisruption": {
		Name: "Phase Disruption",
		Text: "Psionic Storm lowers Physical Armor",
	},
	"TassadarPrismaticLink": {
		Name: "Prismatic Link",
		Text: "Basic Attack hits multiple targets",
	},
	"TassadarPsionicEcho": {
		Name: "Psionic Echo",
		Text: "Psionic Storm may be cast a second time",
	},
	"TassadarPsionicProjection": {
		Name: "Psionic Projection",
		Text: "Increases the range of Plasma Shield, Psionic Storm",
	},
	"TassadarPsionicStormPsiInfusion": {
		Name: "Psi-Infusion",
		Text: "Quest: Psionic Storm returns Mana per target hit",
	},
	"TassadarResonation": {
		Name: "Resonation",
		Text: "Psionic Storm enhances Distortion Beam",
	},
	"TassadarShieldBattery": {
		Name: "Shield Battery",
		Text: "Activate to grant Plasma Shields to all nearby allies",
	},
	"TassadarTemplarsWill": {
		Name: "Templar's Will",
		Text: "Quest: Attacking Heroes empowers Distortion Beam",
	},
	"ThrallAlphaWolf": {
		Name: "Alpha Wolf",
		Text: "Feral Spirit has longer Root, empowers Attacks",
	},
	"ThrallAncestralWrath": {
		Name: "Ancestral Wrath",
		Text: "Activate to damage and steal Health from a Hero",
	},
	"ThrallCrashLightning": {
		Name: "Crash Lightning",
		Text: "Quest: Multiple Hero hits empower Chain Lightning",
	},
	"ThrallEchooftheElements": {
		Name: "Echo of the Elements",
		Text: "Quest: Killing enemies grants Chain Lightning charge",
	},
	"ThrallFeralResilience": {
		Name: "Feral Resilience",
		Text: "Feral Spirit grants Physical Armor, Frostwolf Resilience",
	},
	"ThrallFrostwolfPack": {
		Name: "Frostwolf Pack",
		Text: "Quest: Reduce Feral Spirit cooldown, cost",
	},
	"ThrallHeroicAbilityEarthquake": {
		Name: "Earthquake",
		Text: "Slows enemies in a massive area",
	},
	"ThrallHeroicAbilitySundering": {
		Name: "Sundering",
		Text: "Damages, pushes, and Stuns enemies in a line",
	},
	"ThrallMaelstromWeapon": {
		Name: "Maelstrom Weapon",
		Text: "Quest:  Windfury attacks increase Move Speed",
	},
	"ThrallMasteryEarthenShields": {
		Name: "Earthen Shields",
		Text: "Grants allies Shields in area",
	},
	"ThrallMasteryFrostwolfsGrace": {
		Name: "Frostwolf's Grace",
		Text: "Activate Frostwolf Resilience to instantly heal",
	},
	"ThrallMasteryGraceOfAir": {
		Name: "Grace Of Air",
		Text: "Windfury grants double Frostwolf Resilience",
	},
	"ThrallMasteryManaTide": {
		Name: "Mana Tide",
		Text: "Frostwolf Resilience restores Mana, lowers cooldowns",
	},
	"ThrallMasteryRollingThunder": {
		Name: "Rolling Thunder",
		Text: "Chain Lightning restores Mana, increased range",
	},
	"ThrallMasteryStoneWolves": {
		Name: "Stone Wolves",
		Text: "Increases Feral Spirit root duration",
	},
	"ThrallMasteryTempestFury": {
		Name: "Tempest Fury",
		Text: "Windfury's final strike hits 3 times",
	},
	"ThrallMasteryWorldbreaker": {
		Name: "Worldbreaker",
		Text: "Sundering creates impassable terrain",
	},
	"ThrallSpiritShield": {
		Name: "Spirit Shield",
		Text: "Periodically gain temporary Spell Armor",
	},
	"ThrallThunderstorm": {
		Name: "Thunderstorm",
		Text: "Quest: Hit Heroes to make Chain Lightning Slow",
	},
	"ThrallWindRush": {
		Name: "Wind Rush",
		Text: "Activate to teleport and gain Windfury",
	},
	"ThrallWindStalker": {
		Name: "Wind Stalker",
		Text: "Quest: Windfury attacks empower Windfury",
	},
	"TinkerCombatStyleBreakitDown": {
		Name: "Break it Down!",
		Text: "Increases Salvager cooldown reduction",
	},
	"TinkerCombatStyleClockwerkSteamFists": {
		Name: "Clockwerk Steam Fists",
		Text: "Quest: Basic Attacks increase Rock-It! Turret duration",
	},
	"TinkerCombatStyleReduceReuseRecycle": {
		Name: "Reduce, Reuse, Recycle",
		Text: "Enemy Minions can drop Scrap",
	},
	"TinkerCombatStyleScrapoMaticSmelter": {
		Name: "Scrap-o-Matic Smelter",
		Text: "Increases Salvager Mana restoration",
	},
	"TinkerGoblinRepairs": {
		Name: "Goblin Repairs",
		Text: "Quest: Regen Globes increase Health Regen",
	},
	"TinkerHeroicAbilityGravOBomb3000": {
		Name: "Grav-O-Bomb 3000",
		Text: "Pulls and damages enemies",
	},
	"TinkerHeroicAbilityRoboGoblin": {
		Name: "Robo-Goblin",
		Text: "Increase Attack Damage, Armor, Move Speed",
	},
	"TinkerHiredGoons": {
		Name: "Hired Goons",
		Text: "Empower nearby Mercenaries. Turrets gain Merc Armor",
	},
	"TinkerItsRainingScrap": {
		Name: "It's Raining Scrap",
		Text: "Activate to create Scrap",
	},
	"TinkerMasteryARKReaktor": {
		Name: "ARK Reaktor",
		Text: "Reduces Xplodium Charge cooldown, can give turret charges",
	},
	"TinkerMasteryEZPZDimensionalRipper": {
		Name: "EZ-PZ Dimensional Ripper",
		Text: "Deth Lazor disables enemies",
	},
	"TinkerMasteryEngineGunk": {
		Name: "Engine Gunk",
		Text: "Rock-it! Turrets slow enemies",
	},
	"TinkerMasteryExtraTNT": {
		Name: "Extra TNT",
		Text: "Xplodium Charge deals more damage",
	},
	"TinkerMasteryGoblinFusion": {
		Name: "Goblin Fusion",
		Text: "Deth Lazor deals more damage",
	},
	"TinkerMasteryHyperfocusCoils": {
		Name: "Hyperfocus Coils",
		Text: "Deth Lazor charges faster",
	},
	"TinkerMasteryLongRangedTurrets": {
		Name: "Long-Ranged Turrets",
		Text: "Increases Rock-it! Turret attack range",
	},
	"TinkerMasteryMechaLord": {
		Name: "Mecha-Lord",
		Text: "Gain Armor, Basic Attack damage",
	},
	"TinkerMasteryMiniatureBlackHole": {
		Name: "Miniature Black Hole",
		Text: "Increases Grav-O-Bomb radius and damage",
	},
	"TinkerMasteryQuikReleaseCharge": {
		Name: "Kwik Release Charge",
		Text: "Xplodium Charge gains second charge",
	},
	"TinkerMasteryRockitTurretXL": {
		Name: "Rock-It! Turret XL",
		Text: "Rock-It! Turrets attack multiple enemies",
	},
	"TinkerMasteryTurretStorage": {
		Name: "Turret Storage",
		Text: "Increases Rock-it! Turret charges",
	},
	"TinkerMasteryXTraLargeBombs": {
		Name: "X-Tra Large Bombs",
		Text: "Xplodium Charge casts when Stunned, Rooted",
	},
	"TinkerSuperiorSchematics": {
		Name: "Superior Schematics",
		Text: "Increases Rock-it! Turret range",
	},
	"TinkerTalentFirinMahLazorz": {
		Name: "Firin' Mah Lazorz",
		Text: "Rock-It! Turrets fire Deth Lazorz",
	},
	"TracerBulletSprayMelee": {
		Name: "Bullet Spray",
		Text: "Melee hits all enemies in range",
	},
	"TracerBulletTime": {
		Name: "Bullet Time",
		Text: "Basic Attacks reduce Blink cooldown",
	},
	"TracerCompositionBPulseBomb": {
		Name: "Composition B",
		Text: "Hitting Heroes with Pulse Bomb leaves a second",
	},
	"TracerFocusFire": {
		Name: "Focus Fire",
		Text: "Unloading a magazine on an enemy does bonus damage",
	},
	"TracerGetStuffedMelee": {
		Name: "Get Stuffed!",
		Text: "Reduces Melee cooldown and knocks away",
	},
	"TracerIsthataHealthPack": {
		Name: "Is That a Health Pack?!",
		Text: "Increases Regen Globe and Healing Fountain healing",
	},
	"TracerJumper": {
		Name: "Jumper",
		Text: "Increases Blink charges",
	},
	"TracerLeechingRounds": {
		Name: "Leeching Rounds",
		Text: "Basic Attacks against Heroes heal",
	},
	"TracerLockedandLoaded": {
		Name: "Locked and Loaded",
		Text: "Time Reload perfectly to gain Basic Attack damage",
	},
	"TracerPartingGift": {
		Name: "Parting Gift",
		Text: "Using Recall leaves bombs",
	},
	"TracerPulseRounds": {
		Name: "Pulse Rounds",
		Text: "Increases Pulse Bomb range and charge rate",
	},
	"TracerPulseStrikeMelee": {
		Name: "Pulse Strike",
		Text: "Increases Melee Pulse Bomb charge rate",
	},
	"TracerQuantumSpike": {
		Name: "Quantum Spike",
		Text: "Increases Pulse Bomb single target damage",
	},
	"TracerRicochetHeroWeapon": {
		Name: "Ricochet",
		Text: "Basic Attacks can bounce to a nearby enemy",
	},
	"TracerSleightofHand": {
		Name: "Sleight of Hand",
		Text: "Reload twice as fast",
	},
	"TracerSlipstreamRecall": {
		Name: "Slipstream",
		Text: "Recall rewinds more time",
	},
	"TracerSpatialEcho": {
		Name: "Spatial Echo",
		Text: "Takedowns grant Blink charges",
	},
	"TracerStickyBomb": {
		Name: "Sticky Bomb",
		Text: "Increases Pulse Bomb area and adds slow",
	},
	"TracerTotalRecallRecall": {
		Name: "Total Recall",
		Text: "Recall can restore Health, but has increased cooldown",
	},
	"TracerTracerRounds": {
		Name: "Tracer Rounds",
		Text: "Basic Attacks reveal enemies",
	},
	"TracerUntouchable": {
		Name: "Untouchable",
		Text: "Quest: Takedowns increase Attack Damage",
	},
	"TychusCombatTactician": {
		Name: "Combat Tactician",
		Text: "Basic Attacks reduce Run and Gun cooldown",
	},
	"TychusFullyLoaded": {
		Name: "Fully Loaded",
		Text: "Reduces Minigun cooldown",
	},
	"TychusHeroicAbilityCommandeerOdin": {
		Name: "Commandeer Odin",
		Text: "Call down an Odin to pilot",
	},
	"TychusHeroicAbilityDrakkenLaserDrill": {
		Name: "Drakken Laser Drill",
		Text: "Call down a Laser to attack nearby enemies",
	},
	"TychusInTheRhythm": {
		Name: "In the Rhythm",
		Text: "Quest: Basic Attacks increase Minigun duration",
	},
	"TychusMasterAssassin": {
		Name: "Master Assassin",
		Text: "Quest: Takedowns increase Attack Speed",
	},
	"TychusMasteryDrakkenLaserFocusingDiodes": {
		Name: "Focusing Diodes",
		Text: "Increase Laser Drill range and damage",
	},
	"TychusMasteryFragGrenadeConcussionGrenade": {
		Name: "Concussion Grenade",
		Text: "Increases Frag Grenade radius and knockback",
	},
	"TychusMasteryFragGrenadeTitanGrenade": {
		Name: "Titan Grenade",
		Text: "Increases Frag Grenade damage to Heroes",
	},
	"TychusMasteryLeadRain": {
		Name: "Lead Rain",
		Text: "Overkill slows targets",
	},
	"TychusMasteryOdinBigRedButton": {
		Name: "Big Red Button",
		Text: "Increases Odin duration and fires Nukes",
	},
	"TychusMasteryOverkillArmorPiercingRounds": {
		Name: "Armor Piercing Rounds",
		Text: "Increases Overkill damage to main target",
	},
	"TychusMasteryQuarterback": {
		Name: "Quarterback",
		Text: "Increased Frag Grenade range",
	},
	"TychusMasteryRunandGunDash": {
		Name: "Dash",
		Text: "Quest: Regen Globes empower Run and Gun",
	},
	"TychusMasteryRunandGunStimPack": {
		Name: "Stim Pack",
		Text: "Bonus Movement Speed and Attack Speed",
	},
	"TychusMasteryShredderGrenade": {
		Name: "Shredder Grenade",
		Text: "Increased Frag Grenade radius",
	},
	"TychusMasterySprayNPray": {
		Name: "Spray 'n' Pray",
		Text: "Increases Overkill range",
	},
	"TychusNeosteelCoating": {
		Name: "Neosteel Coating",
		Text: "Activate to gain Spell Armor",
	},
	"TychusPressTheAdvantage": {
		Name: "Press the Advantage",
		Text: "Run and Gun increases Tychus's Basic Attack range",
	},
	"TychusRelentlessSoldier": {
		Name: "Relentless Soldier",
		Text: "Stuns and Roots grant Armor",
	},
	"TychusRunAndGunBobAndWeave": {
		Name: "Bob and Weave",
		Text: "Adds Run and Gun charges and reduces Mana cost",
	},
	"TychusSizzlinAttacks": {
		Name: "Sizzlin' Attacks",
		Text: "Basic Attacks deal more damage to Heroes",
	},
	"TychusThatsTheStuff": {
		Name: "That's the Stuff!",
		Text: "Minigun damage heals Tychus",
	},
	"TychusTheBiggerTheyAre": {
		Name: "The Bigger They Are...",
		Text: "Minigun Damage modified based on Target Health",
	},
	"TyraelHeroicAbilityJudgement": {
		Name: "Judgment",
		Text: "Charge and stun an enemy Hero",
	},
	"TyraelHeroicAbilitySanctification": {
		Name: "Sanctification",
		Text: "Create zone of Invulnerable",
	},
	"TyraelMasteryAngelicAbsorption": {
		Name: "Angelic Absorption",
		Text: "Grants Health Regeneration when attacked",
	},
	"TyraelMasteryAngelicMight": {
		Name: "Angelic Might",
		Text: "Increased Basic Attack damage",
	},
	"TyraelMasteryElDruinsMightAngelsGrace": {
		Name: "Angel's Grace",
		Text: "El'druin's Might teleport grants Movement Speed",
	},
	"TyraelMasteryElDruinsMightBladeOfJustice": {
		Name: "Blade of Justice",
		Text: "El'druin's Might teleport increases Attack Speed",
	},
	"TyraelMasteryElDruinsMightHoradricReforging": {
		Name: "Horadric Reforging",
		Text: "Reduces El'druin's Might cooldown when enemy is hit",
	},
	"TyraelMasteryEvenInDeath": {
		Name: "Even In Death",
		Text: "Can use Basic Abilities before exploding",
	},
	"TyraelMasteryHolyGround": {
		Name: "Holy Ground",
		Text: "Teleporting using El'druin's Might makes a blocking field",
	},
	"TyraelMasteryJudgmentAngelofJustice": {
		Name: "Angel of Justice",
		Text: "Judgment has increased range, reduced cooldown",
	},
	"TyraelMasteryPassiveProtectioninDeath": {
		Name: "Protection in Death",
		Text: "Upon death, grants nearby allies a Shield",
	},
	"TyraelMasteryPurgeEvil": {
		Name: "Purge Evil",
		Text: "Increased Smite damage to Heroes",
	},
	"TyraelMasteryRighteousnessReciprocate": {
		Name: "Reciprocate",
		Text: "Shield explodes on breaking",
	},
	"TyraelMasterySalvation": {
		Name: "Salvation",
		Text: "Stronger Shield per nearby allied Hero",
	},
	"TyraelMasterySanctificationHolyArena": {
		Name: "Holy Arena",
		Text: "Sanctification increased duration, adds damage bonus",
	},
	"TyraelMasterySwiftRetribution": {
		Name: "Swift Retribution",
		Text: "Increases Smite's Movement Speed bonus and duration",
	},
	"TyraelMasteryZealotry": {
		Name: "Zealotry",
		Text: "Increases Righteousness' duration and reduces cooldown",
	},
	"TyrandeCelestialAttunement": {
		Name: "Celestial Attunement",
		Text: "Light of Elune removes Stuns and Silences",
	},
	"TyrandeDarnassianArchery": {
		Name: "Darnassian Archery",
		Text: "Consecutive Basic Attacks deal more damage",
	},
	"TyrandeElunesChosen": {
		Name: "Elune's Chosen",
		Text: "Basic Attacks heal target ally",
	},
	"TyrandeEmpower": {
		Name: "Empower",
		Text: "Hitting Heroes reduces Sentinel cooldown",
	},
	"TyrandeEyesOfTheHuntress": {
		Name: "Eyes of the Huntress",
		Text: "Shadowstalk reveals enemy Heroes",
	},
	"TyrandeHarshMoonlight": {
		Name: "Harsh Moonlight",
		Text: "Sentinel Slows and reduces damage dealt",
	},
	"TyrandeHeroicAbilityShadowstalk": {
		Name: "Shadowstalk",
		Text: "Stealth and gain Movement Speed",
	},
	"TyrandeHeroicAbilityStarfall": {
		Name: "Starfall",
		Text: "Damages and slows enemies in an area",
	},
	"TyrandeHuntressFury": {
		Name: "Huntress' Fury",
		Text: "Hunter's Mark makes Basic Attacks splash",
	},
	"TyrandeIcebladeArrows": {
		Name: "Iceblade Arrows",
		Text: "Gain Attack Speed, reduce enemy damage dealt",
	},
	"TyrandeKaldoreiResistance": {
		Name: "Kaldorei Resistance",
		Text: "Light of Elune grants Spell Armor",
	},
	"TyrandeMarkofMending": {
		Name: "Mark of Mending",
		Text: "Marked enemies heal attackers",
	},
	"TyrandeMasteryLightofEluneQuickeningBlessing": {
		Name: "Quickening Blessing",
		Text: "Light of Elune increases Movement Speed",
	},
	"TyrandeMasteryLunarBlaze": {
		Name: "Lunar Blaze",
		Text: "Quest: Hit Heroes to empower Lunar Flare",
	},
	"TyrandeMasteryLunarFlareShootingStar": {
		Name: "Shooting Star",
		Text: "Lunar Flare deals extra damage and refunds Mana on hit",
	},
	"TyrandeMasterySentinelPierce": {
		Name: "Pierce",
		Text: "Sentinel hits multiple targets",
	},
	"TyrandeMasteryShadowstalkHuntersSwiftness": {
		Name: "Hunter's Swiftness",
		Text: "Allies gain Movement Speed ",
	},
	"TyrandeMasteryStarfallCelestialWrath": {
		Name: "Celestial Wrath",
		Text: "Starfall applies Hunter's Mark to Heroes",
	},
	"TyrandeMoonlitArrows": {
		Name: "Moonlit Arrows",
		Text: "Increase Light of Elune cooldown reduction",
	},
	"TyrandeOverflowingLight": {
		Name: "Overflowing Light",
		Text: "Increase Light of Elune healing at high Health",
	},
	"TyrandeRanger": {
		Name: "Ranger",
		Text: "Quest: Distance increases Sentinel damage",
	},
	"TyrandeRangersMark": {
		Name: "Ranger's Mark",
		Text: "Quest: Basic Attacks empower Hunter's Mark",
	},
	"TyrandeSearingArrows": {
		Name: "Searing Arrows",
		Text: "Activate to increase Basic Attack damage at the cost of Mana",
	},
	"TyrandeShootingStar": {
		Name: "Shooting Star",
		Text: "Basic Attacks periodically cast Lunar Flare",
	},
	"TyrandeTrueshotAura": {
		Name: "Trueshot Aura",
		Text: "Gain Attack Damage aura",
	},
	"UtherBeaconOfLight": {
		Name: "Beacon of Light",
		Text: "Increase Holy Light self-healing",
	},
	"UtherCombatStyleHammerOfTheLightbringer": {
		Name: "Hammer of the Lightbringer",
		Text: "Quest: Reduce Hammer of Justice cooldown",
	},
	"UtherEternalDevotionArmorOfFaith": {
		Name: "Armor of Faith",
		Text: "Holy Light recharges faster while disabled",
	},
	"UtherEternalDevotionDivineProtection": {
		Name: "Divine Protection",
		Text: "Devotion stacks twice",
	},
	"UtherEternalDevotionGuardianOfAncientKings": {
		Name: "Guardian of Ancient Kings",
		Text: "Increase Devotion's Armor on disabled targets",
	},
	"UtherHammerOfJusticePursuitOfJustice": {
		Name: "Pursuit of Justice",
		Text: "Hammer of Justice grants Move Speed",
	},
	"UtherHammerOfJusticeWellMet": {
		Name: "Well Met",
		Text: "Hammer of Justice weakens enemies",
	},
	"UtherHandOfProtection": {
		Name: "Hand of Protection",
		Text: "Make an ally Unstoppable",
	},
	"UtherHardenedFocus": {
		Name: "Hardened Focus",
		Text: "Faster cooldowns near full Health",
	},
	"UtherHeroicAbilityDivineShield": {
		Name: "Divine Shield",
		Text: "Give ally Invulnerable and increased Move Speed",
	},
	"UtherHeroicAbilityDivineStorm": {
		Name: "Divine Storm",
		Text: "Damage and stun nearby enemies",
	},
	"UtherHolyFire": {
		Name: "Holy Fire",
		Text: "Deals damage to nearby enemies",
	},
	"UtherHolyLightSilverTouch": {
		Name: "Silver Touch",
		Text: "Quest: Mitigate damage to improve Holy Light",
	},
	"UtherHolyRadianceTyrsDeliverance": {
		Name: "Tyr's Deliverance",
		Text: "Holy Radiance increases healing received",
	},
	"UtherMasteryBenediction": {
		Name: "Benediction",
		Text: "Reduce Mana cost, cooldown of next ability",
	},
	"UtherMasteryBlessedChampion": {
		Name: "Blessed Champion",
		Text: "Holy Light makes Basic Attacks heal",
	},
	"UtherMasteryBulwarkOfLightDivineShield": {
		Name: "Bulwark of Light",
		Text: "Increase Divine Shield duration, reduce cooldown",
	},
	"UtherMasteryDivineHurricaneDivineStorm": {
		Name: "Divine Hurricane",
		Text: "Increase Divine Storm radius, reduce cooldown",
	},
	"UtherMasteryHolyShock": {
		Name: "Holy Shock",
		Text: "Holy Light can deal damage",
	},
	"UtherMasteryRedemption": {
		Name: "Redemption",
		Text: "Resurrect after Eternal Vanguard ends",
	},
	"UtherMasteryWaveofLightHolyRadiance": {
		Name: "Wave of Light",
		Text: "Quest: Holy Radiance increases Devotion duration",
	},
	"ValeeraAmbushAssassinate": {
		Name: "Assassinate",
		Text: "Ambush deals more damage to isolated Heroes",
	},
	"ValeeraAmbushDeathFromAbove": {
		Name: "Death From Above",
		Text: "Ambush teleports Valeera behind her victim",
	},
	"ValeeraBladeFlurryFatalFinesse": {
		Name: "Fatal Finesse",
		Text: "Quest: Hit Heroes with Blade Flurry",
	},
	"ValeeraCheapShotBlind": {
		Name: "Blind",
		Text: "Cheap Shot also blinds enemies",
	},
	"ValeeraCloakOfShadows": {
		Name: "Cloak of Shadows",
		Text: "Become Unstoppable, gain Spell Armor",
	},
	"ValeeraCloakOfShadowsEnvelopingShadows": {
		Name: "Enveloping Shadows",
		Text: "Vanish grants Cloak of Shadows",
	},
	"ValeeraColdBlood": {
		Name: "Cold Blood",
		Text: "Activate to increase Eviscerate damage",
	},
	"ValeeraCombatReadiness": {
		Name: "Combat Readiness",
		Text: "Spending Combo Points grants Block",
	},
	"ValeeraCripplingPoison": {
		Name: "Crippling Poison",
		Text: "Spell Damage slows enemies",
	},
	"ValeeraEviscerateExposeArmor": {
		Name: "Expose Armor",
		Text: "Eviscerate lowers target's Armor",
	},
	"ValeeraEviscerateSliceAndDice": {
		Name: "Slice and Dice",
		Text: "Eviscerate increases Attack Speed",
	},
	"ValeeraGarroteHemorrhage": {
		Name: "Hemorrhage",
		Text: "Garrote increases Basic Attack damage",
	},
	"ValeeraGarroteRupture": {
		Name: "Rupture",
		Text: "Basic Attacks increase Garrote's damage",
	},
	"ValeeraGarroteStrangle": {
		Name: "Strangle",
		Text: "Garrote reduces Spell Power",
	},
	"ValeeraSinisterStrikeMutilate": {
		Name: "Mutilate",
		Text: "Increases Sinister Strike damage but reduces range",
	},
	"ValeeraSinisterStrikeRelentlessStrikes": {
		Name: "Relentless Strikes",
		Text: "Reduces Sinister Strike Energy cost",
	},
	"ValeeraSinisterStrikeSealFate": {
		Name: "Seal Fate",
		Text: "Empower Sinister Strike damage, Combo Points",
	},
	"ValeeraSmokeBomb": {
		Name: "Smoke Bomb",
		Text: "Create an obscuring cloud of smoke",
	},
	"ValeeraSmokeBombAdrenalineRush": {
		Name: "Adrenaline Rush",
		Text: "Smoke Bomb increases Energy regeneration",
	},
	"ValeeraThistleTea": {
		Name: "Thistle Tea",
		Text: "Activate to instantly restore Energy",
	},
	"ValeeraVanishElusiveness": {
		Name: "Elusiveness",
		Text: "Increases Vanish Movement Speed bonus",
	},
	"ValeeraVanishInitiative": {
		Name: "Initiative",
		Text: "Stealth Abilities award more Combo Points",
	},
	"ValeeraVanishNightslayer": {
		Name: "Nightslayer",
		Text: "Reduces Vanish cooldown",
	},
	"ValeeraVanishSubtlety": {
		Name: "Subtlety",
		Text: "Stealth Abilities increase Energy regeneration",
	},
	"ValeeraVigor": {
		Name: "Vigor",
		Text: "Quest: Regen Globes increase max Energy",
	},
	"ValeeraWoundPoison": {
		Name: "Wound Poison",
		Text: "Spell Damage reduces enemy healing received",
	},
	"VarianBannerOfDalaran": {
		Name: "Banner of Dalaran",
		Text: "Banner that grants Spell Power",
	},
	"VarianBannerOfIronforge": {
		Name: "Banner of Ironforge",
		Text: "Banner that grants Armor",
	},
	"VarianBannerOfStormwind": {
		Name: "Banner of Stormwind",
		Text: "Banner that grants Movement Speed",
	},
	"VarianBannersGloryToTheAlliance": {
		Name: "Glory to the Alliance",
		Text: "Banner also increases healing effects",
	},
	"VarianChargeJuggernaut": {
		Name: "Juggernaut",
		Text: "Increase Charge damage",
	},
	"VarianChargeWarbringer": {
		Name: "Warbringer",
		Text: "Charge has reduced cooldown and Mana cost",
	},
	"VarianColossusSmash": {
		Name: "Colossus Smash",
		Text: "Gain Damage, Lose HealthSmash enemies and lower their Armor",
	},
	"VarianColossusSmashMasterAtArms": {
		Name: "Master at Arms",
		Text: "Colossus Smash hits enemies near the target",
	},
	"VarianDemoralizingShout": {
		Name: "Demoralizing Shout",
		Text: "Reduce damage of nearby Heroes",
	},
	"VarianHighKingsQuestQuest": {
		Name: "High King's Quest",
		Text: "Quest:  Complete Quests to gain Attack Damage ",
	},
	"VarianLionsFangLionsMawQuest": {
		Name: "Lion's Maw",
		Text: "Quest: Hit Heroes to empower Lion's Fang",
	},
	"VarianMortalStrike": {
		Name: "Mortal Strike",
		Text: "Heroic Strike reduces enemy healing",
	},
	"VarianParryLivebytheSword": {
		Name: "Live by the Sword",
		Text: "Parry lasts longer, blocking reduces cooldown",
	},
	"VarianParryOverpower": {
		Name: "Overpower",
		Text: "Parrying a Hero's attack will refresh Heroic Strike",
	},
	"VarianParryShieldWall": {
		Name: "Shield Wall",
		Text: "Parry prevents all damage",
	},
	"VarianSecondWind": {
		Name: "Second Wind",
		Text: "Basic Attacks heal, more when low Health",
	},
	"VarianShatteringThrow": {
		Name: "Shattering Throw",
		Text: "Deal massive damage to enemy Shields",
	},
	"VarianTaunt": {
		Name: "Taunt",
		Text: "Gain Health, ArmorForce a Hero to attack Varian",
	},
	"VarianTauntVigilance": {
		Name: "Vigilance",
		Text: "Incoming attacks reduce Taunt cooldown",
	},
	"VarianTwinBladesOfFury": {
		Name: "Twin Blades of Fury",
		Text: "Gain Attack Speed, Lose DamageBasic Attacks reduce Heroic Strike cooldown",
	},
	"VarianTwinBladesOfFuryFrenzy": {
		Name: "Frenzy",
		Text: "Increase Twin Blades of Fury bonuses",
	},
	"VarianVictoryRush": {
		Name: "Victory Rush",
		Text: "Basic Attacks periodically grant a large heal",
	},
	"WitchDoctorAnnihilatingSpirits": {
		Name: "Annihilating Spirit",
		Text: "Increase range and speed of Ravenous Spirit",
	},
	"WitchDoctorBigVoodoo": {
		Name: "Big Voodoo",
		Text: "Increase Voodoo Ritual bonuses",
	},
	"WitchDoctorBloodRitual": {
		Name: "Blood Ritual",
		Text: "Voodoo Ritual restores Health and Mana",
	},
	"WitchDoctorDeadRush": {
		Name: "Dead Rush",
		Text: "Zombie Wall deals more damage and uproots",
	},
	"WitchDoctorFreshCorpses": {
		Name: "Fresh Corpses",
		Text: "Reduces Zombie Wall cooldown",
	},
	"WitchDoctorGuardianToads": {
		Name: "Guardian Toads",
		Text: "Toads grant Armor",
	},
	"WitchDoctorHeroicAbilityGargantuan": {
		Name: "Gargantuan",
		Text: "Summon a Gargantuan to guard an area",
	},
	"WitchDoctorHeroicAbilityRavenousSpirits": {
		Name: "Ravenous Spirit",
		Text: "Channel a damaging, movable spirit ",
	},
	"WitchDoctorHexedCrawlers": {
		Name: "Hexed Crawlers",
		Text: "Corpse Spiders restore Health and Mana",
	},
	"WitchDoctorHumongoid": {
		Name: "Humongoid",
		Text: "Reduce Gargantuan cooldown and mana cost",
	},
	"WitchDoctorPandemic": {
		Name: "Pandemic",
		Text: "Quest: Empower Plague of Toads",
	},
	"WitchDoctorRingOfPoison": {
		Name: "Ring of Poison",
		Text: "Deals damage in center area",
	},
	"WitchDoctorSoulHarvest": {
		Name: "Soul Harvest",
		Text: "Activate to increase Health and Spell Power",
	},
	"WitchDoctorSpiderColony": {
		Name: "Spider Colony",
		Text: "Corpse Spiders reduce other Basic Ability cooldowns",
	},
	"WitchDoctorSpiritofArachyr": {
		Name: "Spirit of Arachyr",
		Text: "More Corpse Spiders against single enemy",
	},
	"WitchDoctorSuperstition": {
		Name: "Superstition",
		Text: "Gain Spell Armor while not being Attacked",
	},
	"WitchDoctorThingOfTheDeep": {
		Name: "Thing of the Deep",
		Text: "Increase the range of Nazeebo's Abilities",
	},
	"WitchDoctorToadAffinity": {
		Name: "Toad Affinity",
		Text: "Damage against heroes reduces cooldown",
	},
	"WitchDoctorToadsofHugeness": {
		Name: "Toads of Hugeness",
		Text: "Toads become more powerful with each hop",
	},
	"WitchDoctorVileInfection": {
		Name: "Vile Infection",
		Text: "Quest: Empower Voodoo Ritual",
	},
	"WitchDoctorWidowmakers": {
		Name: "Widowmakers",
		Text: "Quest: Empower Corpse Spiders",
	},
	"WizardAetherWalker": {
		Name: "Aether Walker",
		Text: "Teleport costs no Mana when out of combat",
	},
	"WizardArcaneOrbArcaneOrbit": {
		Name: "Arcane Orbit",
		Text: "Increases Arcane Orb range",
	},
	"WizardArcaneOrbTriumvirate": {
		Name: "Triumvirate",
		Text: "Reduces Arcane Orb cooldown at max range",
	},
	"WizardArcaneOrbZeisVengeance": {
		Name: "Zei's Vengeance",
		Text: "Arcane Orb does more damage from afar",
	},
	"WizardArchonPurePower": {
		Name: "Archon: Pure Power",
		Text: "Use Disintegrate repeatedly, but no other Abilities",
	},
	"WizardAstralPresence": {
		Name: "Astral Presence",
		Text: "Increases Mana regen while low Mana",
	},
	"WizardCannoneer": {
		Name: "Cannoneer",
		Text: "Abilities empower Basic Attacks",
	},
	"WizardDisintegrateTemporalFlux": {
		Name: "Temporal Flux",
		Text: "Disintegrate slows enemies",
	},
	"WizardDominance": {
		Name: "Dominance",
		Text: "Takedowns restore Health",
	},
	"WizardFireflies": {
		Name: "Fireflies",
		Text: "Magic Missiles move faster, lower cooldown and Mana cost",
	},
	"WizardForceArmor": {
		Name: "Force Armor",
		Text: "Magic Missiles grants Spell Armor",
	},
	"WizardGlassCannon": {
		Name: "Glass Cannon",
		Text: "Gain Spell Power but lose max Health",
	},
	"WizardHeroicAbilityDisintegrate": {
		Name: "Disintegrate",
		Text: "Long range channeled beam",
	},
	"WizardHeroicAbilityWaveOfForce": {
		Name: "Wave Of Force",
		Text: "Damage and knock enemies back",
	},
	"WizardMagicMissilesChargedBlast": {
		Name: "Charged Blast",
		Text: "Magic Missiles marks enemies for bonus damage",
	},
	"WizardMagicMissilesMirrorball": {
		Name: "Mirrorball",
		Text: "Fires 2 extra Magic Missiles and increases Mana cost",
	},
	"WizardMagicMissilesSeeker": {
		Name: "Seeker",
		Text: "Increases Magic Missiles single-target damage",
	},
	"WizardPowerHungry": {
		Name: "Power Hungry",
		Text: "Regen Globes give Spell Power and more Mana",
	},
	"WizardTalRashasElements": {
		Name: "Tal Rasha's Elements",
		Text: "Cycle abilities to gain Spell Power",
	},
	"WizardTeleportCalamity": {
		Name: "Calamity",
		Text: "Teleport deals damage",
	},
	"WizardTeleportDiamondSkin": {
		Name: "Diamond Skin",
		Text: "Teleport grants a Shield",
	},
	"WizardTeleportIllusionist": {
		Name: "Illusionist",
		Text: "Increases Teleport range, damage refreshes cooldown",
	},
	"WizardWaveOfForceRepulsion": {
		Name: "Repulsion",
		Text: "Increases cast range and knockback distance",
	},
	"ZagaraCombatStyleMedusaBlades": {
		Name: "Medusa Blades",
		Text: "Basic Attacks hit 3 additional targets",
	},
	"ZagaraHeroicAbilityDevouringMaw": {
		Name: "Devouring Maw",
		Text: "Eats enemies in area",
	},
	"ZagaraHeroicAbilityNydusAssault": {
		Name: "Nydus Network",
		Text: "Empowers Creep and can create Nydus Worms",
	},
	"ZagaraMasteryBanelingMassacre": {
		Name: "Baneling Massacre",
		Text: "Increases number of Banelings",
	},
	"ZagaraMasteryBileDrop": {
		Name: "Bile Drop",
		Text: "Quest: Infested Drop impacts increase damage",
	},
	"ZagaraMasteryBroodExpansion": {
		Name: "Brood Expansion",
		Text: "Reduces the cooldown of Hunter Killer",
	},
	"ZagaraMasteryCentrifugalHooks": {
		Name: "Centrifugal Hooks",
		Text: "Banelings can travel twice as far",
	},
	"ZagaraMasteryCorpseFeeders": {
		Name: "Corpse Feeders",
		Text: "Infested Drop cooldown is reduced and Roachlings take less damage",
	},
	"ZagaraMasteryCorrosiveSaliva": {
		Name: "Corrosive Saliva",
		Text: "Empowers Hunter Killers against Heroes",
	},
	"ZagaraMasteryEndlessCreep": {
		Name: "Endless Creep",
		Text: "Increases Creep Tumor's cast range and duration",
	},
	"ZagaraMasteryEnvenomedSpines": {
		Name: "Envenomed Spines",
		Text: "Empower Zagara's next Basic Attack",
	},
	"ZagaraMasteryGroovedSpines": {
		Name: "Grooved Spines",
		Text: "Increased Hunter Killer range and damage",
	},
	"ZagaraMasteryHydraliskTransfusion": {
		Name: "Hydralisk Transfusion",
		Text: "Hunter Killer attacks restore Health",
	},
	"ZagaraMasteryInfest": {
		Name: "Infest",
		Text: "Increases nearby allied Ranged Minion's damage",
	},
	"ZagaraMasteryMetabolicBoost": {
		Name: "Metabolic Boost",
		Text: "Increased Movement Speed bonus on Creep",
	},
	"ZagaraMasteryMutalisk": {
		Name: "Mutalisk",
		Text: "Hunter Killer summons a Mutalisk instead",
	},
	"ZagaraMasteryProtectiveCoating": {
		Name: "Protective Coating",
		Text: "Gain Armor on Creep",
	},
	"ZagaraMasteryReconstitution": {
		Name: "Reconstitution",
		Text: "Increased Health Regeneration on Creep",
	},
	"ZagaraMasterySerratedSpines": {
		Name: "Serrated Spines",
		Text: "Quest: Each Basic Attack against a Hero increases Zagara's Attack Damage by .",
	},
	"ZagaraMasteryTumorClutch": {
		Name: "Tumor Clutch",
		Text: "Decreased Creep Tumor cooldown, Mana cost removed",
	},
	"ZagaraMasteryTyrantMaw": {
		Name: "Tyrant Maw",
		Text: "More damage and cooldown reduction from Takedowns",
	},
	"ZagaraMasteryVolatileAcid": {
		Name: "Volatile Acid",
		Text: "Increase Baneling range and non-Hero damage",
	},
	"ZagaraViscousAcid": {
		Name: "Viscous Acid",
		Text: "Banelings also slow enemies",
	},
	"ZaryaDefensiveShielding": {
		Name: "Defensive Shielding",
		Text: "Personal Barrier, Shield Ally grant Physical Armor",
	},
	"ZaryaEnergyBornInBattle": {
		Name: "Born in Battle",
		Text: "Reduces cooldowns at high Energy",
	},
	"ZaryaEnergyEnduranceTraining": {
		Name: "Endurance Training",
		Text: "Gain Armor at high Energy",
	},
	"ZaryaEnergyHitMe": {
		Name: "Hit Me",
		Text: "Shield damage absorbed contributes more Energy",
	},
	"ZaryaEnergyMaximumChargeQuest": {
		Name: "Maximum Charge",
		Text: "Regen Globes grant EnergyQuest: Increases maximum Energy",
	},
	"ZaryaExpulsionZoneClearOut": {
		Name: "Clear Out",
		Text: "Increases Expulsion Zone radius and provides Energy",
	},
	"ZaryaGravitonSurgeGravityKills": {
		Name: "Gravity Kills",
		Text: "Graviton Surge has increased duration and deals damage",
	},
	"ZaryaHeroicAbilityExpulsionZone": {
		Name: "Expulsion Zone",
		Text: "Launch a bomb that continually pushes enemies",
	},
	"ZaryaHeroicAbilityGravitonSurge": {
		Name: "Graviton Surge",
		Text: "Launch a gravity bomb that draws in enemies",
	},
	"ZaryaPainIsTemporary": {
		Name: "Pain is Temporary",
		Text: "Consume Energy to gain a Shield",
	},
	"ZaryaParticleGrenadeDemolitionsExpertQuest": {
		Name: "Demolitions Expert",
		Text: "Quest: Increases Particle Grenade radius and reduces cooldown",
	},
	"ZaryaParticleGrenadeGrenadier": {
		Name: "Grenadier",
		Text: "Particle Grenade charges are returned at once",
	},
	"ZaryaParticleGrenadePinpointAccuracy": {
		Name: "Pinpoint Accuracy",
		Text: "Particle Grenade deals more damage in center area",
	},
	"ZaryaParticleGrenadePlasmaShock": {
		Name: "Plasma Shock",
		Text: "Particle Grenade slows enemies near the center",
	},
	"ZaryaPersonalBarrierExplosiveBarrier": {
		Name: "Explosive Barrier",
		Text: "Personal Barrier deals area damage",
	},
	"ZaryaPersonalBarrierIAmTheStrongest": {
		Name: "I Am the Strongest",
		Text: "Increases Personal Barrier Shield amount",
	},
	"ZaryaPersonalBarrierSpellBarrier": {
		Name: "Spell Barrier",
		Text: "Personal Barrier grants Spell Armor",
	},
	"ZaryaPersonalBarrierUnstoppableCompetitor": {
		Name: "Unstoppable Competitor",
		Text: "Personal Barrier grants Unstoppable",
	},
	"ZaryaShieldAllyCleansingShield": {
		Name: "Cleansing Shield",
		Text: "Shield Ally removes all disabling effects",
	},
	"ZaryaShieldAllyGainTrain": {
		Name: "Gain Train",
		Text: "Shield Ally now affects a second ally",
	},
	"ZaryaShieldAllyGiveMeTwentyQuest": {
		Name: "Give Me Twenty",
		Text: "Quest: Increases Shield Ally absorb amount and reduces cooldown",
	},
	"ZaryaShieldAllySpeedBarrier": {
		Name: "Speed Barrier",
		Text: "Shield Ally increases allied Hero Movement Speed",
	},
	"ZaryaShieldAllyTogetherWeAreStrong": {
		Name: "Together We Are Strong",
		Text: "Allied damage dealt with Shield Ally contributes to Energy",
	},
	"ZaryaUnyieldingDefender": {
		Name: "Unyielding Defender",
		Text: "Activate to reset Shield cooldowns",
	},
	"ZaryaWeaponFeelTheHeat": {
		Name: "Feel the Heat",
		Text: "Basic Attack deals more damage to close enemies",
	},
	"ZaryaWeaponToTheLimit": {
		Name: "To the Limit",
		Text: "Increases Basic Attack size at high Energy",
	},
	"ZeratulComboSlash": {
		Name: "Combo Slash",
		Text: "Abilities increase Basic Attack damage",
	},
	"ZeratulGiftoftheXelNaga": {
		Name: "Gift of the Xel’Naga",
		Text: "Allies unaffected by Void Prison, enemies slowed",
	},
	"ZeratulGrimTask": {
		Name: "Grim Task",
		Text: "Quest: Hero Takedowns increase Spell Power by %, up to % . This bonus Spell Power is lost on death.",
	},
	"ZeratulHeroicAbilityShadowAssault": {
		Name: "Shadow Assault",
		Text: "Basic Attacks teleport to enemies",
	},
	"ZeratulHeroicAbilityVoidPrison": {
		Name: "Void Prison",
		Text: "Time Stop targets in an area",
	},
	"ZeratulMasterWarpBlade": {
		Name: "Master Warp-Blade",
		Text: "Bonus damage every third Basic Attack",
	},
	"ZeratulMasteryGreaterCleaveCleave": {
		Name: "Greater Cleave",
		Text: "Increases radius of Cleave",
	},
	"ZeratulMendingStrikes": {
		Name: "Mending Strikes",
		Text: "Basic Attacks heal",
	},
	"ZeratulNerazimFury": {
		Name: "Nerazim Fury",
		Text: "Grants Life Steal and increased duration",
	},
	"ZeratulRendingCleave": {
		Name: "Rending Cleave",
		Text: "Cleave deals additional damage over time",
	},
	"ZeratulSeekerInTheDark": {
		Name: "Seeker in the Dark",
		Text: "Teleport to Singularity Spike target",
	},
	"ZeratulSentencedtoDeath": {
		Name: "Sentenced to Death",
		Text: "Deal bonus damage to marked targets",
	},
	"ZeratulShadowHunter": {
		Name: "Shadow Hunter",
		Text: "Quest: Gather Regen Globes to empower Blink",
	},
	"ZeratulShroudofAdun": {
		Name: "Shroud of Adun",
		Text: "Gain a shield while under Permanent Cloak",
	},
	"ZeratulSlipintoShadow": {
		Name: "Slip into Shadow",
		Text: "Blink has 2 charges, but higher cooldown",
	},
	"ZeratulVoidSlash": {
		Name: "Void Slash",
		Text: "Cleave empowered against multiple Heroes",
	},
	"ZeratulVorpalBlade": {
		Name: "Vorpal Blade",
		Text: "Teleport to last attacked enemy",
	},
	"ZeratulWormhole": {
		Name: "Wormhole",
		Text: "Can return from Blink location",
	},
	"ZuljinAmaniRageTalent": {
		Name: "Amani Rage",
		Text: "Activate to lose Health and restore it over time",
	},
	"ZuljinAmaniResilience": {
		Name: "Amani Resilience",
		Text: "Taz'dingo duration increased, restores Health",
	},
	"ZuljinArcaniteAxes": {
		Name: "Arcanite Axes",
		Text: "Quest: Increase Twin Cleave damage, lower cooldown",
	},
	"ZuljinBoneslicer": {
		Name: "Boneslicer",
		Text: "Empower Grievous Throw",
	},
	"ZuljinBuzzsaw": {
		Name: "Buzzsaw",
		Text: "Guillotine keeps traveling forward",
	},
	"ZuljinEnsnare": {
		Name: "Ensnare",
		Text: "Throw a net that Roots enemies",
	},
	"ZuljinEyeOfZuljin": {
		Name: "Eye of Zul'jin",
		Text: "Basic Attacks grant Move Speed",
	},
	"ZuljinFerocity": {
		Name: "Ferocity",
		Text: "Increases Berserker's Attack Speed bonus",
	},
	"ZuljinForestMedicine": {
		Name: "Forest Medicine",
		Text: "Regeneration cannot be interrupted",
	},
	"ZuljinGuillotine": {
		Name: "Guillotine",
		Text: "Deal heavy area damage",
	},
	"ZuljinHeadhunter": {
		Name: "Headhunter",
		Text: "Quest: Takedowns increase damage",
	},
	"ZuljinLacerate": {
		Name: "Lacerate",
		Text: "Increase Twin Cleave Slow",
	},
	"ZuljinLetTheKillingBegin": {
		Name: "Let the Killing Begin",
		Text: "Killing enemies increases Attack Speed",
	},
	"ZuljinNoMercyTalent": {
		Name: "No Mercy!",
		Text: "Basic Attacks ignore Armor",
	},
	"ZuljinRecklessness": {
		Name: "Recklessness",
		Text: "Low Health increases damage",
	},
	"ZuljinSwirlingDeath": {
		Name: "Swirling Death",
		Text: "Twin Cleave Axes keep spinning",
	},
	"ZuljinTazdingo": {
		Name: "Taz'dingo!",
		Text: "Become unkillable",
	},
	"ZuljinTrollsBlood": {
		Name: "Troll's Blood",
		Text: "Regeneration heals more",
	},
	"ZuljinViciousAssault": {
		Name: "Vicious Assault",
		Text: "Increase Grievous Throw damage bonus",
	},
	"ZuljinVoodooShuffle": {
		Name: "Voodoo Shuffle",
		Text: "Activate to remove Roots and Slows",
	},
	"ZuljinWrongPlaceWrongTime": {
		Name: "Wrong Place Wrong Time",
		Text: "Bonus Twin Cleave damage at apex",
	},
	"ZuljinYouWantAxe": {
		Name: "You Want Axe?",
		Text: "Quest: Every  Basic Attacks against Heroes permanently increases Basic Attack damage by .Reward: After attacking Heroes  times, Basic Attack range is increased by .Reward: After attacking Heroes  times, Twin Cleave now revolves twice.",
	},
}
