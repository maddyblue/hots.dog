import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Fetch, pct, toLength } from './common';
import SortedTable from './SortedTable';

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
		const name = this.state.name;
		this.setState({ loading: true });
		Fetch(
			'/api/get-player-by-name?name=' + encodeURIComponent(this.state.name),
			data => {
				if (this.state.name === name) {
					this.setState({ ids: data, loading: false });
				}
			}
		);
	}
	render() {
		let names;
		if (this.state.loading) {
			names = 'loading...';
		} else if (!this.state.ids) {
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
		if (this.state.Games.length) {
			game = (
				<SortedTable
					name="player"
					sort="Date"
					headers={[
						{
							name: 'Build',
							header: 'Patch',
							desc: true,
						},
						{
							name: 'Date',
							cell: v => new Date(v).toLocaleString(),
							desc: true,
						},
						{
							name: 'Hero',
							cell: v => <Link to={'/talents/' + encodeURI(v)}>{v}</Link>,
						},
						{
							name: 'HeroLevel',
							header: 'Level',
							title: 'hero level',
							desc: true,
						},
						{
							name: 'Winner',
							header: 'Won',
							cell: v => (v ? 'win' : 'loss'),
							desc: true,
						},
						{
							name: 'Length',
							title: 'game length',
							cell: toLength,
						},
						{
							name: 'Map',
						},
						{
							name: 'Mode',
							header: 'Game Mode',
							cell: v => this.props.Modes[v],
						},
					]}
					data={this.state.Games}
				/>
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
