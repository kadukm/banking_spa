import React from 'react';


export default class Company extends React.Component {
    render() {
        return (
            <div>{this.props.match.params.companyId}</div>
        )
    }
}
