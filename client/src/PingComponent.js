import React, { Component } from "react";
import axios from "axios";

class PingComponent extends Component {
    state = {
        delegationLists: []
    };

    componentWillMount() {
        axios
            .get("lists")
            .then(res => {
                this.setState({ delegationLists: res.data });
            })
            .catch(function(error) {
                console.log(error);
            });
    }

    render() {
        return (
            <ul>
                {this.state.delegationLists.map(list => (
                    <li>{list.name}</li>
                ))}
            </ul>
        );
    }
}

export default PingComponent;
