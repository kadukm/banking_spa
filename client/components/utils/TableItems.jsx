import React from 'react';

export function TableKey({ name }) {
    return (
        <span className="table__key table__cell">
            {name}
        </span>
    )
}

export function TableCell(props) {
    return (
        <span className="table__cell">
            {props.children}
        </span>
    )
}
