import * as React from "react";
import { Box } from "./Box";

export class EnterRepo extends React.Component<{ loadRepo: (org: string, name: string) => void }, { repo: string }> {
    constructor (props: any) {
        super(props);

        this.state = {
            repo: "",
        };
    }

    setRepo = (e: React.FormEvent<HTMLButtonElement>) => {
        this.setState({
            repo: e.currentTarget.value,
        });
    };

    loadRepo = (e: any) => {
        const [org, name] = this.state.repo.split('/')
        this.props.loadRepo(org, name);
    };

    checkEnter = (e: React.KeyboardEvent<HTMLButtonElement>) => {
        if (e.key == 'Enter') this.loadRepo(null);
    }

    render = () => {
        return (
            <div id="enterRepo">
                <Box>
                    <div id="enterRepoMain">
                        <h3 id="enterRepoTitle">Enter a Go repo to test:</h3>
                        <div>
                            <b>github.com/ </b>
                            <input id="enterRepoInput" onChange={this.setRepo} onKeyDown={this.checkEnter} type="text"></input>
                        </div>
                    </div>
                    <button id="load" onClick={this.loadRepo}>LOAD</button>
                </Box>
            </div>
        );
    }
}
