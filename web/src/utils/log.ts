import { H } from 'highlight.run';

const Log = H

Log.init('', {
	serviceName: "frontend-app",
	tracingOrigins: true,
	networkRecording: {
		enabled: true,
		recordHeadersAndBody: true,
		urlBlocklist: [
			// insert full or partial urls that you don't want to record here
			// Out of the box, Highlight will not record these URLs (they can be safely removed):
			"https://www.googleapis.com/identitytoolkit",
			"https://securetoken.googleapis.com",
		],
	},
});

export default Log