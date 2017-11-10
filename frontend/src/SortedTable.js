// @flow

import React, { Component } from 'react';
import { createCookie, readCookie } from './common';

const tableSortings = {};
const sortTableCookie = 'sort-table-';
function setTableSort(name, sort, dir) {
	const val = sort + ',' + (dir === true).toString();
	createCookie(sortTableCookie + name, val);
	const sortings = tableSortings[name];
	const st = {
		sort: sort,
		sortDir: dir,
	};
	sortings.forEach(v => {
		v.setState(st);
	});
}
function registerTableSort(that) {
	const name = that.props.name;
	if (!tableSortings[name]) {
		tableSortings[name] = [];
	}
	tableSortings[name].push(that);
}
function unregisterTableSort(that) {
	const name = that.props.name;
	tableSortings[name] = tableSortings[name].filter(v => v !== that);
}

type Cmp = (any, any) => number;

type Header = {
	name: string,
	header?: string,
	cell?: any => string,
	desc?: boolean,
	title?: string,
	cmp?: Cmp,
};

type Headers = Header[];

type Props = {
	name: string,
	headers: Headers,
	sort: string,
	data: any[],
};

type State = {
	sort: string,
	sortDir: boolean,
	lookup: { [string]: any },
};

type Lookup = {
	[string]: {
		name: string,
		header: string,
		cell: any => string,
		desc: boolean,
		title?: string,
		cmp: Cmp,
	},
};

class SortedTable extends Component<Props, State> {
	constructor(props: Props) {
		super(props);
		const lookup = this.lookup(props.headers);
		let sort = props.sort;
		if (!this.props.name || !lookup[props.sort]) {
			debugger;
		}
		let dir = lookup[props.sort].desc === true;
		const cookie = readCookie(sortTableCookie + this.props.name);
		registerTableSort(this);
		if (cookie) {
			const sp = cookie.split(',', 2);
			if (sp.length === 2) {
				sort = sp[0];
				dir = sp[1] === 'true';
			}
		}
		this.state = {
			lookup: lookup,
			sort: sort,
			sortDir: dir,
		};
	}
	componentWillUnmount() {
		unregisterTableSort(this);
	}
	componentWillUpdate(nextProps: Props, nextState: State, nextContext: any) {
		nextState.lookup = this.lookup(nextProps.headers);
	}
	lookup(headers: Headers): Lookup {
		const lookup: Lookup = {};
		headers.forEach(h => {
			if (!h.cell) {
				h.cell = v => v;
			}
			if (!h.header) {
				h.header = h.name;
			}
			lookup[h.name] = {
				name: h.name,
				header: h.header || h.name,
				cell: h.cell || (v => v),
				desc: h.desc === true,
				title: h.title,
				cmp: h.cmp || defaultCmp,
			};
		});
		return lookup;
	}
	sort = (sort: string) => {
		if (this.state.lookup[sort].cmp === null) {
			return;
		}
		let dir;
		if (this.state.sort === sort) {
			dir = !this.state.sortDir;
		} else {
			dir = this.state.lookup[sort].desc === true;
		}
		setTableSort(this.props.name, sort, dir);
	};
	sortClass = (col: string) => {
		if (col !== this.state.sort) {
			return '';
		}
		const dir = this.state.sortDir ? 'desc' : 'asc';
		return 'sort-' + dir;
	};
	render() {
		const data = this.props.data;
		data.sort((a: any, b: any): number =>
			this.state.lookup[this.state.sort].cmp(
				a[this.state.sort],
				b[this.state.sort]
			)
		);
		if (this.state.sortDir) {
			data.reverse();
		}
		const body = data.map((row, i) => (
			<tr key={i}>
				{this.props.headers.map(h => (
					<td key={h.name}>
						{this.state.lookup[h.name].cell(row[h.name], row)}
					</td>
				))}
			</tr>
		));
		const inner = [
			<thead key="thead">
				<tr>
					{this.props.headers.map(h => {
						let { name, title, header } = h;
						return (
							<th
								key={name}
								onClick={() => this.sort(name)}
								className={this.sortClass(name)}
								title={title}
							>
								{header}
							</th>
						);
					})}
				</tr>
			</thead>,
			<tbody key="tbody">{body}</tbody>,
		];
		if (this.props.notable) {
			return inner;
		}
		return <table className="sorted">{inner}</table>;
	}
}

const defaultCmp: Cmp = (a, b): number => {
	switch (typeof a) {
		case 'number':
		case 'boolean':
			return a - b;
		case 'string':
			return a.localeCompare(b);
		default:
			if (a instanceof Date) {
				return a - b;
			} else {
				debugger;
				return 0;
			}
	}
};

export default SortedTable;
