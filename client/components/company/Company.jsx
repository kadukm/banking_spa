import React from 'react';
import Header from './parts/Header.jsx'
import Body from './parts/Body.jsx'
import Footer from './parts/Footer.jsx'
import Delimiter from '../utils/Delimiter.jsx'
import apiBaseUrl from '../../config.js' 

function getDataFrom(url) {
    return new Promise(
        (resolve, reject) => fetch(url)
            .then(response => response.json())
            .then(response => {
                if (response.ok)
                    return resolve(response.result)
                else
                    throw new Error(response.result)
            })
            .catch(err => reject(err))
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
            .then(([companyInfo]) => {
                this.setState({ok: true, companyInfo: companyInfo})
            })
            .catch(err => {this.setState({ok: false}); console.log(err)})
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
                <Footer {...this.state.companyInfo}/>
            </div>
        )
    }
}
