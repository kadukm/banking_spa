import React from 'react';
import Header from './parts/Header.jsx'
import Body from './parts/Body.jsx'
import Footer from './parts/Footer.jsx'
import Delimiter from '../utils/Delimiter.jsx'
import apiBaseUrl from '../../config.js' 
import '../../styles/Company.css'

const maxProductsCount = 4

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
        this.state = {
            ok: undefined,
            companyInfo: undefined,
            products: undefined}
    }

    componentDidMount() {
        const companyID = this.props.match.params.companyID
        Promise.all([
                getDataFrom(`${apiBaseUrl}/api/companies/${companyID}`),
                getDataFrom(`${apiBaseUrl}/api/companies/${companyID}/products?maxcount=${maxProductsCount}`)
            ])
            .then(([companyInfo, products]) => {
                this.setState({ok: true, companyInfo: companyInfo, products: products})
            })
            .catch(err => {this.setState({ok: false}); console.log(err)})
    }

    render() {
        if (this.state.ok === undefined)
            return <div>Loading...</div>
        if (this.state.ok === false)
            return <div>Can't load data about {this.props.match.params.companyID}</div>
        return (
            <div className="company-page">
                <Header {...this.state.companyInfo}/>
                <Delimiter />
                <Body {...this.state.companyInfo}/>
                <Delimiter />
                <Footer {...this.state.companyInfo} products={this.state.products}/>
            </div>
        )
    }
}
