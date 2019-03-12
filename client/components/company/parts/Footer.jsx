import React from 'react';

export default class Footer extends React.Component {
    render() {
        return (
            <section className="box-wrapper">
                <h3>О компании {this.props.status} {this.props.name}</h3>
                <div>{this.props.info}</div>
                <a href={this.props.full_info_path}>Полная информация</a>
            </section>
        )
    }
}