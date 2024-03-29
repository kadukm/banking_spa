import React from 'react';
import { BrowserRouter, Switch, Route } from "react-router-dom";
import AdminPanel from './admin-panel/AdminPanel.jsx'
import Company from './company/Company.jsx'
import Login from './Login.jsx'
import '../styles/App.css'

export default class App extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <Switch>
                    <Route
                        path="/admin-panel"
                        component={AdminPanel} />
                    <Route
                        path="/companies/:companyID"
                        component={Company}
                    />
                    <Route
                        path="/login"
                        component={Login}
                    />
                    <Route
                        component={() => <div>Wrong URL!</div>}
                    />
                </Switch>
            </BrowserRouter>
        );
    }
}
