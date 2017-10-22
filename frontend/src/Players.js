import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Fetch, pct, toLength } from './common';

class Players extends Component {
	constructor(props) {
		super(props);
		this.state = {
			name: '',
			ids: [],
		};
		this.update = this.update.bind(this);
		this.onSubmit = this.onSubmit.bind(this);
	}
	update(ev) {
		const st = {};
		st[ev.target.name] = ev.target.value;
		this.setState(st);
	}
	onSubmit(ev) {
		ev.preventDefault();
		if (!this.state.name) {
			return;
		}
		Fetch(
			'/api/get-player-by-name?name=' + encodeURIComponent(this.state.name),
			data => this.setState({ ids: data })
		);
	}
	render() {
		let names;
		if (!this.state.ids) {
			names = 'No matches.';
		} else {
			names = this.state.ids.map(e => (
				<li key={e.ID}>
					<Link to={'/players/' + e.ID}>{e.Name}</Link>
				</li>
			));
		}
		return (
			<div>
				<form onSubmit={this.onSubmit}>
					<label>Search for player by name</label>
					<input
						type="text"
						name="name"
						value={this.state.name}
						onChange={this.update}
					/>
					<input className="button-primary" type="submit" value="search" />
				</form>
				<ul>{names}</ul>
			</div>
		);
	}
}

class Player extends Component {
	componentDidMount() {
		Fetch(
			'/api/get-player-data?id=' +
				encodeURIComponent(this.props.match.params.id),
			data => {
				if (!data.Skills) {
					data.Skills = [];
				}
				if (!data.Games) {
					data.Games = [];
				}
				this.setState(data);
			}
		);
	}
	percentile(s) {
		const modes = this.props.BuildStats[s.Build];
		if (!modes) {
			return;
		}
		const stats = modes[s.Mode];
		if (!stats) {
			return;
		}
		const qs = Object.keys(stats.Quantile);
		qs.sort((a, b) => a - b);
		if (s.Skill < stats.Quantile[qs[0]]) {
			return pct(0, 0);
		}
		for (let i = 1; i < qs.length; i++) {
			const q1 = qs[i];
			const p1 = stats.Quantile[q1];
			if (p1 >= s.Skill) {
				const q0 = qs[i - 1];
				const p0 = stats.Quantile[q0];
				// If p lies a fraction f of the way from p{i} to p{i+1}, define the pth
				// quantile to be:
				// Q(p) = (1-f)Q(p{i}) + fQ(p{i+1})
				const f = (s.Skill - p0) / (p1 - p0);
				const qp = (1 - f) * q0 + f * q1;
				return pct(qp, 0);
			}
		}
		return pct(100, 0);
	}
	render() {
		if (!this.state) {
			return 'loading...';
		}
		const skills = this.state.Skills.map((s, i) => (
			<tr key={i}>
				<td>{s.Build}</td>
				<td>{this.props.Modes[s.Mode]}</td>
				<td>{s.Skill}</td>
				<td>{this.percentile(s)}</td>
			</tr>
		));
		const games = this.state.Games.map((g, i) => (
			<tr key={i}>
				<td>{g.Build}</td>
				<td>{new Date(g.Date).toLocaleString()}</td>
				<td>
					<Link to={'/talents/' + encodeURI(g.Hero)}>{g.Hero}</Link>
				</td>
				<td>{g.HeroLevel}</td>
				<td>{g.Winner ? 'win' : 'loss'}</td>
				<td>{toLength(g.Length)}</td>
				<td>{g.Map}</td>
				<td>{this.props.Modes[g.Mode]}</td>
			</tr>
		));
		let game, skill;
		if (skills.length) {
			skill = (
				<div>
					<p>Skill rating at the end of each patch with played games:</p>
					<table>
						<thead>
							<tr>
								<th>Patch</th>
								<th>Game Mode</th>
								<th>Skill Rating</th>
								<th>Percentile</th>
							</tr>
						</thead>
						<tbody>{skills}</tbody>
					</table>
				</div>
			);
		}
		if (games.length) {
			game = (
				<table>
					<thead>
						<tr>
							<th>Patch</th>
							<th>Date</th>
							<th>Hero</th>
							<th title="Hero level">Level</th>
							<th>Won</th>
							<th title="Game length">Length</th>
							<th>Map</th>
							<th>Game Mode</th>
						</tr>
					</thead>
					<tbody>{games}</tbody>
				</table>
			);
		}
		return (
			<div>
				<h2>{this.state.Battletag}</h2>
				{skill}
				{game}
			</div>
		);
	}
}

export { Players, Player };
