import * as React from "react";

export class Box extends React.Component {
    constructor (props: any) {
        super(props);
    }

    render = () => {
        return <div className="box">{this.props.children}</div>
    }
}
