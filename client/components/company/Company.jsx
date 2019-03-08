import React from 'react';
import Header from './parts/Header.jsx'
import Body from './parts/Body.jsx'
import Footer from './parts/Footer.jsx'
import Delimiter from '../utils/Delimiter.jsx'
import apiBaseUrl from '../../config.js' 

function getDataFrom(url) {
    return new Promise(
        resolve => fetch(url)
            .then(response => response.json())
            .then(result => resolve(result))
    )
}

export default class Company extends React.Component {
    constructor(props) {
        super(props)
        this.state = {ok: undefined, companyInfo: undefined}
    }

    componentDidMount() {
        const companyID = this.props.match.params.companyID
        Promise.all([getDataFrom(`${apiBaseUrl}/api/companies/${companyID}`)])
            .then(([companyInfo]) => this.setState({ok: true, companyInfo: companyInfo}))
    }

    render() {
        if (this.state.ok === undefined)
            return <div>Loading...</div>
        if (this.state.ok === false)
            return <div>Can't load data about {this.props.match.params.companyID}</div>
        return (
            <div>
                <Header {...this.state.companyInfo}/>
                <Delimiter />
                <Body />
                <Delimiter />
                <Footer />
            </div>
        )
    }
}
