import React from 'react';
import {BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from 'src/components/Home.js'
import Task from 'src/components/Task.js'

function App() {
  return (
  <Router>
    <Route exact path="/" component={Home} />
    <Route path="/task/" component={Task} />
  </Router>
  );
}

export default App;
