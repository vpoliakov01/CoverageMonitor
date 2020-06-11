import * as React from "react";
import { FileInfo, Coverage, FileCoverage } from "./Repo";

export class FileView extends React.Component<{ info: FileInfo, coverage: FileCoverage }, {}> {
    constructor(props: any) {
        super(props);
    }

    render = () => {
        const lines = atob(this.props.info.content).split('\n');
        const isCovered: boolean[] = new Array(lines.length).fill(false);
        if (this.props.coverage) {
            for (const block of this.props.coverage.coverage_blocks) {
                for (let i = block.start_line; i < block.end_line; i++) {
                    isCovered[i] = true;
                }
            }
        }
        const getClassName = (i: number): string => {
            if (!this.props.coverage) return '';
            return isCovered[i] ? 'lineCovered' : 'lineNotCovered';
        }

        return (
            <div>
                <table>
                    <tbody>
                        {
                            lines.map((line, i) => {
                                return (
                                    <tr>
                                        <td className="lineIndex">
                                            <span>{i + 1}</span>
                                        </td>
                                        <td className={getClassName(i)}>
                                            <span className="line">
                                                {line}
                                            </span>
                                        </td>
                                    </tr>
                                );
                            })
                        }
                    </tbody>
                </table>
            </div>
        );
    }
}
