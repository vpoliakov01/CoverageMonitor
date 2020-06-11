import * as React from "react";
import { Box } from "./Box";
import { FileSelector } from "./FileSelector";
import { FileView } from "./FileView";
import { request } from "../utils";

export interface FileInfo {
    path: string,
    content: string,
}

export interface RepoMeta {
    watchers: number,
    language: string,
}

export interface RepoInfo {
    name: string,
    org: string,
    meta: RepoMeta,
    files : FileInfo[],
}

export interface CoverageBlock {
    start_line: number,
    end_line: number,
}

export interface FileCoverage {
    file_name : string,
    coverage_blocks: CoverageBlock[],
}

export interface Coverage {
    files : FileCoverage[],
}

export class Repo extends React.Component<{ info: RepoInfo }, { file: number, coverage: Coverage }> {
    private coverageMap = new Map<string, FileCoverage>(); 
    constructor(props: any) {
        super(props);

        this.state = {
            file: 0,
            coverage: null,
        };
    }

    selectFile = (file: number) => {
        this.setState({
            file,
        })
    }

    loadCoverage = (e: any) => {
        request('POST', `/api/${this.props.info.org}/${this.props.info.name}/test`, (data: any) => {
            if (!data || data.error) throw 'request failed';

            // create a map of covered files
            for (const file of data.result.files) {
                if (file.file_name.startsWith(`github.com/`)) {
                    // remove the repo prefix
                    file.file_name = file.file_name.split('/').slice(3).join('/');
                }
                this.coverageMap.set(file.file_name, file)
            }

            this.setState({
                coverage: data.result,
            });
        });
    }

    render = () => {
        const fileCoverage: FileCoverage = this.coverageMap.get(this.props.info.files[this.state.file].path);
        return (
            <div>
                <div id="repoHeader">
                    <div id="repoTitle">
                        <span id="repoName">{this.props.info.org}/{this.props.info.name}</span>
                        <span id="star">★</span>
                        <b>{
                            this.props.info.meta.watchers >= 1000 ?
                                `${Math.round(this.props.info.meta.watchers / 100) / 10}k` :
                                `${this.props.info.meta.watchers}`
                        }</b>
                        </div>{
                            !this.state.coverage ?
                                <button id="runTests" onClick={this.loadCoverage}>RUN TESTS</button> :
                                <div id="runTestsClicked">TESTS RAN SUCCESSFULLY!</div>
                        }</div>
                <Box>
                    <table>
                        <tbody>
                            <tr>
                                <th id="fileSelectorHeader">
                                    <span>Files</span>
                                    <span id="languageTag">{this.props.info.meta.language.toUpperCase()}</span>
                                </th>
                                <th id="fileViewHeader">
                                    {this.props.info.files[this.state.file].path}
                                </th>
                            </tr>
                            <tr>
                                <td id="fileSelector">
                                    <FileSelector files={this.props.info.files} selectFile={this.selectFile}></FileSelector>
                                </td>
                                <td id="fileView">
                                    <FileView info={this.props.info.files[this.state.file]} coverage={fileCoverage}></FileView>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </Box>
            </div>
        );
    }
}
