import React from 'react';
import { BrowserRouter, Switch, Route } from "react-router-dom";
import '../styles/App.css'

export default class App extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <Switch>
                    <Route
                        path="/admin-panel"
                        component={() => <div>Hello, Admin!</div>} />
                    <Route
                        path="/companies/:companyId"
                        component={({ match }) => <div>Info about {match.params.companyId}!</div>}
                    />
                    <Route
                        component={() => <div>Default</div>}
                    />
                </Switch>
            </BrowserRouter>
        );
    }
}
