import React, { Component } from "react";

class HeaderComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  render() {
    return (
      <div>
        <header className="header">
          <nav>
            <strong>USER MANAGEMENT APP</strong>
          </nav>
        </header>
      </div>
    );
  }
}

export default HeaderComponent;
