import React from 'react';
import { BrowserRouter, Switch, Route } from "react-router-dom";
import AdminPanel from './AdminPanel.jsx'
import Company from './Company.jsx'
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
                        path="/companies/:companyId"
                        component={Company}
                    />
                </Switch>
            </BrowserRouter>
        );
    }
}
