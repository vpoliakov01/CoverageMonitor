import * as React from "react";
import { EnterRepo } from "./EnterRepo";
import { Repo, RepoInfo } from "./Repo";
import { request } from "../utils";

interface AppState {
    repoInfo: RepoInfo,
    loaded: boolean,
}

export class App extends React.Component<{}, AppState> {
    constructor (props: any) {
        super(props);
    }

    loadRepo = (org: string, name: string) => {
        request('GET', `/api/${org}/${name}/info`, (data: any) => {
            if (!data || data.error) throw 'request failed';

            this.setState({
                repoInfo: data.result,
                loaded: true,
            });
        });
    }

    render = () => {
        return (
            <div id="app">{
                this.state && this.state.loaded ?
                    <Repo info={this.state.repoInfo}></Repo> :
                    <EnterRepo loadRepo={this.loadRepo}></EnterRepo>
            }</div>
        );
    }
}
