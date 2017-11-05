import React from 'react';
import { Link } from 'react-router-dom';

const skillPercentiles = [0, 30, 50, 70, 90, 95, 100];

const BuildsOpts = props => {
	let builds = props.builds.map(b => (
		<option key={b.ID} value={b.ID}>
			{b.ID} ({new Date(b.Start).toLocaleDateString()} -{' '}
			{new Date(b.Finish).toLocaleDateString()})
		</option>
	));
	builds.unshift(
		<option key="latest" value="">
			latest ({props.builds[0].ID})
		</option>
	);
	return builds;
};

const Filter = props => {
	let maps = props.Maps.map(m => <option key={m}>{m}</option>);
	maps.unshift(
		<option key="" value="">
			All Maps
		</option>
	);
	let modeKeys = Object.keys(props.Modes);
	modeKeys.sort().reverse();
	let modes = modeKeys.map(k => (
		<option key={k} value={k}>
			{props.Modes[k]}
		</option>
	));
	modes.unshift(
		<option key="" value="">
			All Game Modes
		</option>
	);
	let heroLevels = [1, 5, 10, 20].map(v => <option key={v}>{v}</option>);
	/*
	const skills = skillPercentiles.map(v => (
		<option key={v} value={v}>
			{v + 'th'}
		</option>
	));
	const build = props.build === 'latest' ? props.Builds[0].ID : props.build;
	const buildStats = props.BuildStats[build];
	const modeStats = buildStats && buildStats[props.mode];
	const disableSkill = !modeStats || !props.mode;
	let skillTitle = 'Enabled when a game mode is selected.';
	if (!buildStats || (!modeStats && props.mode)) {
		skillTitle = 'Skill ratings not yet calculated.';
	}
	*/
	return (
		<div>
			<div className="row">
				<div className="column">
					<label>Map</label>
					<select name="map" value={props.map} onChange={props.handleChange}>
						{maps}
					</select>
				</div>
				<div className="column">
					<label>Patch</label>
					<select
						name="build"
						value={props.build}
						onChange={props.handleChange}
					>
						<BuildsOpts builds={props.Builds} />
					</select>
				</div>
			</div>
			<div className="row">
				<div className="column">
					<label>Game Mode</label>
					<select name="mode" value={props.mode} onChange={props.handleChange}>
						{modes}
					</select>
				</div>
				<div className="column">
					<label>Minimum Hero Level</label>
					<select
						name="herolevel"
						value={props.herolevel}
						onChange={props.handleChange}
					>
						{heroLevels}
					</select>
				</div>
			</div>
			{/*
			<div className="row">
				<div className="column">
					<label title={skillTitle}>Skill Percentile from</label>
					<select
						name="skill_low"
						value={props.skill_low}
						onChange={props.handleChange}
						disabled={disableSkill}
					>
						{skills.slice(0, -1)}
					</select>
				</div>
				<div className="column">
					<label>to</label>
					<select
						name="skill_high"
						value={props.skill_high}
						onChange={props.handleChange}
						disabled={disableSkill}
					>
						{skills.slice(1)}
					</select>
				</div>
			</div>
			*/}
		</div>
	);
};

const HeroHeader = props => {
	const hero = props.heroes.find(e => e.Name === props.name);
	const encHero = encodeURI(props.name);
	const heroSearch = props.build
		? '?build=' + encodeURIComponent(props.build)
		: '';
	const getClass = s =>
		props.history.location.pathname.startsWith('/' + s)
			? 'button button-outline'
			: 'button';
	return (
		<div>
			<h2>
				{hero.Name}{' '}
				<img
					src={'/img/hero_full/' + hero.Slug + '.png'}
					alt={hero.Name}
					style={{ height: '3.4rem' }}
				/>
			</h2>
			<Link
				key="relative"
				className={getClass('heroes')}
				to={'/heroes/' + encHero + heroSearch}
			>
				relative winrates
			</Link>{' '}
			<Link
				key="talents"
				className={getClass('talents')}
				to={'/talents/' + encHero + props.history.location.search}
			>
				talents
			</Link>{' '}
			<Link
				key="compare"
				className={getClass('compare')}
				to={'/compare/' + encHero + props.history.location.search}
			>
				hero comparison
			</Link>
		</div>
	);
};

function Fetch(url, success, error) {
	if (!error) {
		error = alert;
	}
	fetch(url)
		.then(resp => {
			if (resp.status !== 200) {
				resp.text().then(error, error);
				return;
			}
			resp.json().then(success, error);
		})
		.catch(error);
}

function pct(x, n = 1) {
	return x.toFixed(n) + '%';
}

function toLength(l) {
	const mins = Math.trunc(l / 60);
	let secs = (l % 60).toString();
	if (secs.length < 2) {
		secs = '0' + secs;
	}
	return mins + ':' + secs;
}

function createCookie(name, value) {
	localStorage.setItem(name, value);
}

function readCookie(name) {
	return localStorage.getItem(name);
}

export {
	BuildsOpts,
	createCookie,
	Fetch,
	Filter,
	HeroHeader,
	pct,
	readCookie,
	skillPercentiles,
	toLength,
};
