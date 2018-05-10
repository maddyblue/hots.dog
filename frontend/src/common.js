// @flow

import React from 'react';
import { Link } from 'react-router-dom';
import { Helmet } from 'react-helmet';

const skillPercentiles = [0, 7, 42, 77, 92, 99, 100];

const skillBands = [
	'Bronze',
	'Silver',
	'Gold',
	'Platinum',
	'Diamond',
	'Master',
];

const BuildsOpts = (props: { builds: any, dates?: number[] }) => {
	let builds = props.builds.map(b => (
		<option key={b.ID} value={b.ID}>
			{b.ID} ({new Date(b.Start).toLocaleDateString()} -{' '}
			{new Date(b.Finish).toLocaleDateString()})
		</option>
	));
	if (props.dates) {
		const len = props.dates ? props.dates.length : 0;
		props.dates.forEach((d, i) => {
			builds.unshift(
				<option key={d} value={i === len - 1 ? '' : d}>
					last {d} days
				</option>
			);
		});
	} else {
		builds.unshift(
			<option key="latest" value="">
				latest ({props.builds[0].ID})
			</option>
		);
	}
	return builds;
};

const GameModes = (props: { Modes: {}, noAll?: boolean }) => {
	let modeKeys = Object.keys(props.Modes);
	modeKeys.sort().reverse();
	let modes = modeKeys.map(k => (
		<option key={k} value={k}>
			{props.Modes[k]}
		</option>
	));
	if (!props.noAll) {
		modes.unshift(
			<option key="" value="">
				All Game Modes
			</option>
		);
	}
	return modes;
};

const Filter = (props: {
	handleChange: Event => void,
	map: string,
	build: string,
	mode: string,
	herolevel: string,
	Builds: any,
	Maps: string[],
	Modes: {},
	BuildStats: any,
	skill_low: any,
	skill_high: any,
}) => {
	let maps = props.Maps.map(m => <option key={m}>{m}</option>);
	maps.unshift(
		<option key="" value="">
			All Maps
		</option>
	);
	let heroLevels = [1, 5, 10, 20].map(v => <option key={v}>{v}</option>);
	const skills = skillBands.map((v, i) => (
		<option key={v} value={i}>
			{v}
		</option>
	));
	// I'm not sure why the !props.build is needed here, but it seems to be.
	const build =
		!props.build || props.build === 'latest' ? props.Builds[0].ID : props.build;
	const buildStats = props.BuildStats[build];
	const modeStats = buildStats && buildStats[props.mode];
	const disableSkill = !modeStats || !props.mode;
	let skillTitle = 'Enabled when a game mode is selected.';
	if (!buildStats || (!modeStats && props.mode)) {
		skillTitle = 'Skill ratings not yet calculated for this patch.';
	}
	const skillTooltip = disableSkill ? (
		<span className="tip">{skillTitle}</span>
	) : null;
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
						<GameModes Modes={props.Modes} />
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
			<div className={'row' + (disableSkill ? ' tooltip disabled' : '')}>
				{skillTooltip}
				<div className="column">
					<label>Skill band from</label>
					<select
						name="skill_low"
						value={props.skill_low}
						onChange={props.handleChange}
						disabled={disableSkill}
					>
						{skills}
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
						{skills}
					</select>
				</div>
			</div>
		</div>
	);
};

const HeroHeader = (props: {
	heroes: any[],
	name: string,
	build?: string,
	history: any,
	title: any,
	prefix: any,
}) => {
	const hero = props.heroes.find(e => e.Name === props.name);
	if (!hero) {
		return 'unknown hero';
	}
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
			<Title
				{...props.title}
				prefix={props.name + ' ' + props.prefix}
				h={[
					<img
						key="img"
						src={'/img/hero_full/' + hero.Slug + '.png'}
						alt={hero.Name}
						style={{ height: '3.4rem' }}
					/>,
					' ',
				]}
			/>
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

const HeroImg = (props: { name: string, slug: string, link?: string }) => (
	<Link to={'/talents/' + encodeURI(props.name) + (props.link || '')}>
		<img
			src={'/img/hero/' + props.slug + '.png'}
			alt={props.name}
			style={{
				width: '40px',
				height: '40px',
				verticalAlign: 'middle',
				marginRight: '1em',
			}}
		/>
		{props.name}
	</Link>
);

function Fetch(url: string, success: any => void, error?: any => void) {
	if (!error) {
		error = console.log;
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

function pct(x: number, n: number = 1) {
	if (typeof n !== 'number') {
		n = 1;
	}
	return x.toFixed(n) + '%';
}

function toLength(l: number) {
	const mins = Math.trunc(l / 60);
	let secs = (l % 60).toString();
	if (secs.length < 2) {
		secs = '0' + secs;
	}
	return mins + ':' + secs;
}

function createCookie(name: string, value: string) {
	localStorage.setItem(name, value);
}

function readCookie(name: string) {
	return localStorage.getItem(name);
}

function toDate(s: string) {
	return new Date(s).toLocaleString();
}

const TalentImg = (props: { name: string, text?: boolean, data: any }) => {
	let { Name, Text } = props.data;
	let desc;
	if (!Name) {
		const words = props.name.match(/[A-Z][a-z]+/g) || ['', props.name];
		words.shift();
		Name = words.join(' ');
		Text = Name;
	} else if (!props.text) {
		Text = Name + ': ' + Text;
	}
	if (props.text) {
		desc = Name;
	}
	const right = props.text ? '1em' : '2px';
	return (
		<span className="tooltip">
			<img
				key="img"
				src={'/img/talent/' + props.name + '.png'}
				alt={Name}
				style={{
					verticalAlign: 'middle',
					marginRight: right,
					height: '40px',
					width: '40px',
				}}
			/>
			{desc}
			<span className="tip" style={{ whiteSpace: 'pre-line' }}>
				{Text}
			</span>
		</span>
	);
};

const Title = (props: any) => {
	const title =
		props.prefix +
		(props.region ? ' from ' + regionNames[props.region] : '') +
		(props.build ? ' at ' + props.build : '') +
		(props.mode ? ' in ' + props.Modes[props.mode] : '') +
		(props.map ? ' on ' + props.map : '') +
		(props.herolevel !== '5' ? ' above level ' + props.herolevel : '') +
		(props.skill_low > 0 ? ' from ' + skillBands[props.skill_low] : '') +
		(props.skill_high < skillBands.length - 1
			? ' to ' + skillBands[props.skill_high]
			: '');
	return [
		<h2 key="h2">
			{props.h}
			{title}
		</h2>,
		<Helmet key="helmet">
			<title>{title}</title>
		</Helmet>,
	];
};

const regionNames = {
	'1': 'Americas',
	'2': 'Europe',
	'3': 'Asia',
	'5': 'China',
};

const regionCookie = 'region';

export {
	BuildsOpts,
	createCookie,
	Fetch,
	Filter,
	GameModes,
	HeroHeader,
	HeroImg,
	pct,
	readCookie,
	regionCookie,
	regionNames,
	skillPercentiles,
	skillBands,
	toLength,
	toDate,
	TalentImg,
	Title,
};
