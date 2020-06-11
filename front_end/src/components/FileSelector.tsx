import * as React from "react";
import { FileInfo } from "./Repo";

export class FileSelector extends React.Component<{ files: FileInfo[], selectFile: (file: number) => void }, { file: number }> {
    constructor(props: any) {
        super(props);

        this.state = {
            file: 0,
        };
    }

    selectFile = (file: number) => {
        return (e: React.FormEvent<HTMLButtonElement>) => {
            this.props.selectFile(file);
        };
    }

    render = () => {
        return (
            <div id="fileSelector">
                {this.props.files.map((file, i) => {
                    return (
                        <button className="passiveButton" key={`file_${i}`} onClick={this.selectFile(i)}>
                            <FileIcon></FileIcon>
                            {file.path}
                        </button>
                    );
                })}
            </div>
        );
    }
}

const FileIcon = () => {
    return (
        <svg className="fileIcon" aria-label="file" viewBox="0 0 12 16" version="1.1" width="12" height="16" role="img"><path fillRule="evenodd" d="M6 5H2V4h4v1zM2 8h7V7H2v1zm0 2h7V9H2v1zm0 2h7v-1H2v1zm10-7.5V14c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V2c0-.55.45-1 1-1h7.5L12 4.5zM11 5L8 2H1v12h10V5z"></path></svg>
    );
}
