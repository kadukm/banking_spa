import React from 'react';

export default class Footer extends React.Component {
    render() {
        return (
            <section>
                <h3>О компании {this.props.status} {this.props.name}</h3>
                {this.props.products.map(product => <Product {...product} key={product.name}/>)}
                <div>{this.props.info}</div>
                <a href={this.props.full_info_path}>Полная информация</a>
            </section>
        )
    }
}

function Product(props) {
    return (
        <div>
            <img src={props.image_path} alt="product_image" />
            <div>{props.name}</div>
            <div>{props.price}</div>
        </div>
    )

}
        