import React from "react";
import Home from "./home";
import Literature from "./literature";

class Main extends React.Component {

    renderActiveComponent() {
        switch(this.props.activeComponent) {
            case 'home':
                return (<Home />);
            case 'statistics':
                return (<div />)
            case 'literature':
                return (<Literature />);
            default:
                return (<Home />);
        }
    }

    render() {
        return (
            <div className="main">
                {this.renderActiveComponent()}
            </div>
        );
    }
}

export default Main;