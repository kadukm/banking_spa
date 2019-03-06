import React from 'react';
import Header from './parts/Header.jsx'
import Body from './parts/Body.jsx'
import Footer from './parts/Footer.jsx'
import Delimiter from '../utils/Delimiter.jsx'


export default class Company extends React.Component {
    render() {
        return (
            <div>
                <Header companyID={this.props.match.params.companyID}/>
                <Delimiter />
                <Body />
                <Delimiter />
                <Footer />
            </div>
        )
    }
}
