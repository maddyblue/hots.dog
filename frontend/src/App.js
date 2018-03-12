// @flow

import React, { Component } from 'react';
import { Link, Route, withRouter } from 'react-router-dom';
import { Helmet } from 'react-helmet';
import * as usp from 'url-search-params';

import About from './About';
import CompareHero from './Compare';
import { createCookie, Fetch, skillBands } from './common';
import Hero from './Hero';
import HeroWinrates from './HeroWinrates';
import { Players, Player, Game, PlayerGames, PlayerMatchups } from './Players';
import TalentWinrates from './Talents';

import './normalize.css';
import './App.css';
import './milligram/milligram.css';

import logo from './img/icon.svg';

type State = {
	init?: boolean,
	herolevel?: string,
	skill_low?: string,
	skill_high?: string,
	Modes?: any,
	BuildStats?: any,
};

class HotsApp extends Component<{ location: Location, history: any }, State> {
	params: string[];
	defaultHeroLevel: string;
	defaultSkillLow: string;
	defaultSkillHigh: string;

	constructor(props) {
		super(props);
		this.params = [
			'build',
			'map',
			'mode',
			'herolevel',
			'skill_low',
			'skill_high',
		];
		this.defaultHeroLevel = '5';
		this.defaultSkillLow = '0';
		this.defaultSkillHigh = (skillBands.length - 1).toString();
		this.state = this.searchState();
	}
	componentDidMount() {
		Fetch('/api/init', data => {
			data.init = true;
			this.setState(data);
		});
	}
	componentDidUpdate(prevProps, prevState) {
		// This function attempts to sink the state with the URL when the URL changes,
		// for use with back/forward navigation actions.
		const locationChanged = this.props.location !== prevProps.location;
		if (!locationChanged) {
			return;
		}
		var st = this.searchState();
		var updates = {};
		this.params.forEach(key => {
			if (st[key] !== this.state[key]) {
				updates[key] = st[key] || '';
			}
		});
		if (Object.keys(updates).length === 0) {
			return;
		}
		this.setState(updates);
	}
	searchState(): State {
		const st: State = {};
		const search = new usp(window.location.search);
		this.params.forEach(key => {
			if (!search.has(key)) {
				return;
			}
			st[key] = search.get(key);
		});
		if (!st.herolevel) {
			st.herolevel = this.defaultHeroLevel;
		}
		if (!st.skill_low) {
			st.skill_low = this.defaultSkillLow;
		}
		if (!st.skill_high) {
			st.skill_high = this.defaultSkillHigh;
		}
		return st;
	}
	getSearch() {
		const params = [];
		this.params.forEach(key => {
			const value = this.state[key];
			if (!value) {
				return;
			}
			if (key === 'herolevel' && value === this.defaultHeroLevel) {
				return;
			}
			if (
				key === 'skill_low' &&
				(value === this.defaultSkillLow || !this.state.mode)
			) {
				return;
			}
			if (
				key === 'skill_high' &&
				(value === this.defaultSkillHigh || !this.state.mode)
			) {
				return;
			}
			params.push(key + '=' + encodeURIComponent(value));
		});
		return '?' + params.join('&');
	}
	handleChange = event => {
		const st = {};
		st[event.target.name] = event.target.value;
		const ev = +event.target.value;
		switch (event.target.name) {
			case 'skill_low':
				if (ev >= +this.state.skill_high) {
					st['skill_high'] = ev.toString();
				}
				break;
			case 'skill_high':
				if (ev <= +this.state.skill_low) {
					st['skill_low'] = ev.toString();
				}
				break;
			default:
				break;
		}
		this.setState(st, () => {
			const params = this.getSearch();
			createCookie('params', params);
			this.props.history.push({ search: params });
		});
	};
	render() {
		if (!this.state.init) {
			return 'loading...';
		}
		return (
			<main className="wrapper">
				<nav className="navigation">
					<section className="container">
						<Link to="/" className="navigation-title">
							<img
								src={logo}
								alt="hots.dog"
								style={{
									verticalAlign: 'middle',
									width: '30px',
									height: '30px',
								}}
							/>
						</Link>
						<ul className="navigation-list float-right">
							<li className="navigation-item">
								<Link className="navigation-title" to="/players">
									players
								</Link>
							</li>
							<li className="navigation-item">
								<Link className="navigation-link" to="/about">
									about
								</Link>
							</li>
						</ul>
					</section>
				</nav>
				<Helmet>
					<title>hots.dog</title>
					<meta
						name="description"
						content="Heroes of the Storm statistics from aggregating replays."
					/>
				</Helmet>
				<section className="container">
					<Route
						exact
						path="/"
						render={props => (
							<HeroWinrates
								params={this.params}
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
					<Route exact path="/about" component={About} />
					<Route exact path="/players" component={Players} />
					<Route
						exact
						path="/players/:region/:id"
						render={props => (
							<Player
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
					<Route
						exact
						path="/players/:region/:id/games"
						render={props => (
							<PlayerGames
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
					<Route
						exact
						path="/players/:region/:id/matchups"
						render={props => (
							<PlayerMatchups
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
					<Route
						exact
						path="/games/:id"
						render={props => <Game {...props} Modes={this.state.Modes} />}
					/>
					<Route
						exact
						path="/heroes/:hero"
						render={props => (
							<Hero
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
					<Route
						exact
						path="/talents/:hero"
						render={props => (
							<TalentWinrates
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
					<Route
						exact
						path="/compare/:hero"
						render={props => (
							<CompareHero
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
				</section>
			</main>
		);
	}
}

const App = withRouter(HotsApp);

export default App;
