import React from 'react';
import './index.css';

class Header extends React.Component {
    headerButtonIds = ['home', 'statistics', 'literature', 'competition', 'status', 'achieve', 'helper', 'login'];
    stages = [5, 6, 7, 8, 9, 10];
    render = () => {
        return (
        <div className="header">
            <img src="./Media/Logo/BrainZone.png" className="brain-zone-logo" alt="Brain Zone"/>
            <div className="navbar">
                <nav className="header-navbar">
                    { this.headerButtonIds.map((buttonId) => {
                        const className = `header-button ${buttonId}`;
                        const src = `./Media/Header/${buttonId}.png`;
                        const handleButtonClick = () => this.props.setActiveComponent(buttonId);
                        return (
                            <span className={className}>
                                <img className="header-icon" src={src} alt={buttonId} onClick={handleButtonClick}/>
                            </span>
                        );
                    })}
                </nav>
                <nav className="stage-navbar">
                { this.stages.map((stage) => {
                        const className = `button-stage`;
                        const isActive = this.props.activeStage >= stage;
                        const src = `./Media/Header/stage${(isActive ? 'Active' : 'Inactive') + stage}.png`;
                        const handleButtonClick = () => this.props.setActiveStage(stage);
                        return (
                            <span className={className}>
                                <img src={src} alt={stage} onClick={handleButtonClick}/>
                            </span>
                        );
                    })}
                </nav>
            </div>
        </div>
    )}
}

export default Header;