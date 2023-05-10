import React from "react";
import ReactDOM from "react-dom";
import AppFooter from "./AppFooter";
import "./index.css";

function App() {
	return (
		<div className="app">
			<div>
				<h1> Hello, world!</h1>
			</div>

			<AppFooter />
		</div>
	);
}

ReactDOM.render(<App />, document.getElementById("root"));

// It seems that I can use bootstrap with React too...