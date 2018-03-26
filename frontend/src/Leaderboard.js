// @flow

import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import {
	Fetch,
	BuildsOpts,
	regionNames,
	readCookie,
	regionCookie,
	GameModes,
	Title,
} from './common';
import SortedTable from './SortedTable';

type Props = {
	handleChange: any,
	region: string,
	mode: string,
	build?: string,
	Builds: any,
	Modes: any,
};

class Leaderboard extends Component<
	Props,
	{
		Players: any[],
		search: string,
	}
> {
	state = {
		Players: [],
		search: '',
	};
	constructor(props: Props) {
		super(props);
		if (!props.region) {
			this.props.handleChange(
				{
					target: { name: 'region', value: readCookie(regionCookie) || '1' },
				},
				true
			);
		}
		if (!props.mode) {
			this.props.handleChange(
				{
					target: { name: 'mode', value: '3' },
				},
				true
			);
		}
	}
	componentDidMount() {
		this.update();
	}
	componentDidUpdate() {
		this.update();
	}
	makeSearch = () => {
		if (!this.props.mode || !this.props.region) {
			return;
		}
		const build = this.props.build || this.props.Builds[0].ID;
		return (
			'/api/get-leaderboard?build=' +
			build +
			'&mode=' +
			this.props.mode +
			'&region=' +
			this.props.region
		);
	};
	update = () => {
		const search = this.makeSearch();
		if (!search) {
			return;
		}
		if (this.state.search === search) {
			return;
		}
		this.setState({ Players: [], search: search });
		Fetch(search, data => {
			if (search === this.makeSearch()) {
				this.setState(data);
			}
		});
	};
	render() {
		let content;
		if (!this.state.Players.length) {
			content = 'loading...';
		} else {
			content = (
				<SortedTable
					name="leaderboard"
					sort="Rank"
					headers={[
						{
							name: 'Rank',
						},
						{
							name: 'Battletag',
							cell: (v, row) => (
								<Link
									to={
										'/players/' +
										this.props.region +
										'/' +
										row.Blizzid +
										'?build=' +
										(this.props.build || this.props.Builds[0].ID)
									}
								>
									{v}
								</Link>
							),
						},
						{
							name: 'Skill',
							cell: v => v.toFixed(6),
						},
					]}
					data={this.state.Players}
				/>
			);
		}
		return (
			<div>
				<Title {...this.props} prefix="Leaderboard" />
				<div className="row">
					<div className="column">
						<label>Region</label>
						<select
							name="region"
							value={this.props.region}
							onChange={this.props.handleChange}
						>
							{Object.keys(regionNames)
								.sort()
								.map(k => (
									<option key={k} value={k}>
										{regionNames[k]}
									</option>
								))}
						</select>
					</div>
					<div className="column">
						<label>Game Mode</label>
						<select
							name="mode"
							value={this.props.mode}
							onChange={this.props.handleChange}
						>
							<GameModes Modes={this.props.Modes} noAll={true} />
						</select>
					</div>
					<div className="column">
						<label>Patch</label>
						<select
							name="build"
							value={this.props.build}
							onChange={this.props.handleChange}
						>
							<BuildsOpts builds={this.props.Builds} />
						</select>
					</div>
				</div>
				{content}
			</div>
		);
	}
}

export { Leaderboard };
