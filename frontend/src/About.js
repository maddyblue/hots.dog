import React from 'react';
import { Helmet } from 'react-helmet';

const About = props => {
	return (
		<div>
			<Helmet>
				<title>About hots.dog</title>
			</Helmet>
			<h2>about</h2>
			<p>
				This site is a{' '}
				<a href="http://us.battle.net/heroes/en/">Heroes of the Storm</a> game
				aggregator. It fetches data from{' '}
				<a href="http://hotsapi.net/">HotsApi</a> and displays hero winrates
				with various filter options.
			</p>
			<p>You can send feedback by:</p>
			<ul>
				<li>
					emailing <a href="mailto:human@hots.dog">human@hots.dog</a>
				</li>
				<li>
					tweeting <a href="https://twitter.com/mjibson">@mjibson</a>
				</li>
				<li>
					opening a{' '}
					<a href="https://github.com/mjibson/hots.dog/issues/new">
						GitHub issue
					</a>
				</li>
			</ul>
			<p>Our goals are:</p>
			<ul>
				<li>a fast, clean web experience</li>
				<li>
					URL reflects current view; this means you can bookmark a specific
					filter view
				</li>
				<li>open API for others to use</li>
			</ul>
			<p>Our future goals are:</p>
			<ul>
				<li>player MMR calculation</li>
			</ul>
			<p>
				The code is free on GitHub at{' '}
				<a href="https://github.com/mjibson/hots.dog">
					github.com/mjibson/hots.dog
				</a>
				. Technical details:
			</p>
			<ul>
				<li>
					backend is <a href="https://golang.org/">Go</a>
				</li>
				<li>
					database is <a href="https://www.cockroachlabs.com/">CockroachDB</a>
				</li>
				<li>
					frontend is <a href="https://facebook.github.io/react/">React</a>
				</li>
				<li>
					site is deployed with <a href="https://kubernetes.io/">Kubernetes</a>{' '}
					on{' '}
					<a href="https://cloud.google.com/container-engine/">
						Google Container Engine
					</a>
				</li>
			</ul>
		</div>
	);
};

export default About;
