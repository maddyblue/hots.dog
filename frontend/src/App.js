import React, { Component } from 'react';
import { Link, Route, withRouter } from 'react-router-dom';
import * as usp from 'url-search-params';

import About from './About';
import { createCookie, Fetch, skillPercentiles } from './common';
import Hero from './Hero';
import HeroWinrates from './HeroWinrates';
import { Players, Player } from './Players';
import TalentWinrates from './Talents';

import './normalize.css';
import './App.css';
import './milligram.css';

class HotsApp extends Component {
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
		this.defaultSkillHigh = '100';
		this.state = this.searchState();
		this.handleChange = this.handleChange.bind(this);
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
		var locationChanged = this.props.location !== prevProps.location;
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
	searchState() {
		const st = {};
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
	handleChange(event) {
		const st = {};
		st[event.target.name] = event.target.value;
		const ev = +event.target.value;
		switch (event.target.name) {
			case 'skill_low':
				if (ev >= +this.state.skill_high) {
					st['skill_high'] = skillPercentiles[skillPercentiles.indexOf(ev) + 1];
				}
				break;
			case 'skill_high':
				if (ev <= +this.state.skill_low) {
					st['skill_low'] = skillPercentiles[skillPercentiles.indexOf(ev) - 1];
				}
				break;
			default:
				break;
		}
		this.setState(st, () => {
			const params = this.getSearch();
			createCookie('params', params, 50);
			this.props.history.push({ search: params });
		});
	}
	render() {
		if (!this.state.init) {
			return 'loading...';
		}
		return (
			<main className="wrapper">
				<nav className="navigation">
					<section className="container">
						<Link to="/" className="navigation-title">
							home
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
						path="/players/:id"
						render={props => (
							<Player
								{...props}
								Modes={this.state.Modes}
								BuildStats={this.state.BuildStats}
							/>
						)}
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
				</section>
			</main>
		);
	}
}

const App = withRouter(HotsApp);

export default App;
