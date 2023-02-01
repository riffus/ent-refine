import * as RA  from "@pankod/refine-antd";
import * as types  from "./interfaces";
import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';
import CodeEditor from '@uiw/react-textarea-code-editor';

export interface ViewProps<T> {
    value: T | undefined
}

// region Date
export const ER_DateViewOnShow: React.FC<ViewProps<Date>> = ({value}) => {
    return value ? <RA.DateField format="LLL" value={ value }/> : <RA.TextField value="Never"/>
}
export const ER_DateViewOnList: React.FC<ViewProps<Date>> = ER_DateViewOnShow
export const ER_DateViewOnForm: React.FC = (props) => {
    return <RA.DatePicker {...props} showTime={true} />
}
// endregion Date

// region Boolean
export const ER_BooleanViewOnShow: React.FC<ViewProps<types.ER_Boolean>> = ({value}) => {
    return <RA.TextField value={ value ? "Yes" : "No"}/>
}
export const ER_BooleanViewOnList = ER_BooleanViewOnShow
export const ER_BooleanViewOnForm: React.FC = (props)=>{
    return <RA.Radio.Group {...props} >
        <RA.Radio value={true}>Yes</RA.Radio>
        <RA.Radio value={false}>No</RA.Radio>
    </RA.Radio.Group>
}
// endregion Boolean

// region String
export const ER_StringViewOnShow: React.FC<ViewProps<types.ER_String>> = ({value}) => {
    return <RA.Typography.Text copyable={true}>{ value }</RA.Typography.Text>
}
export const ER_StringViewOnList: React.FC<ViewProps<types.ER_String>> = ({value}) => {
    return <RA.Tooltip title={value}>
        <RA.TextField value={value} ellipsis={true} style={ {width: '100px'} }/>
    </RA.Tooltip>
}
export const ER_StringViewOnForm: React.FC = (props) => {
    return <RA.Input {...props} />
}
// endregion String

// region Number
export const ER_NumberViewOnShow: React.FC<ViewProps<types.ER_Number>> = ({value, ...props}) => {
    return <ER_StringViewOnShow value={String(value)} {...props} />
}
export const ER_NumberViewOnList: React.FC<ViewProps<types.ER_Number>> = ({value, ...props}) => {
    return <ER_NumberViewOnShow value={value} {...props} />
}
export const ER_NumberViewOnForm: React.FC = (props) => {
    return <RA.InputNumber {...props} />
}
// endregion Number

// region String List
export const ER_StringListViewOnShow: React.FC<ViewProps<types.ER_StringList>> = ({value}) => {
    return <>{ value?.map( (v, i) => <ER_StringViewOnShow value={ String(i+1) + '. ' + v}/> ) }</>;
}
export const ER_StringListViewOnList: React.FC<ViewProps<types.ER_StringList>> = ({value}) => {
    return <RA.Typography.Text>{ String( value?.length || 0 )+" items" }</RA.Typography.Text>
}
export const ER_StringListViewOnForm: React.FC = (props) => {
    return <RA.Select {...props}  mode="tags" />
}
// endregion

// region Image
export const ER_ImageViewOnShow: React.FC<ViewProps<types.ER_Image>> = ({value}) => {
    return <RA.Image style={ {width: '100%', maxWidth: '256px' } } src={value}/>
}
export const ER_ImageViewOnList: React.FC<ViewProps<types.ER_Image>> = ({value}) => {
    return <RA.Image width={48} src={value}/>
}
export const ER_ImageViewOnForm: React.FC = ER_StringViewOnForm
// endregion Image

// region UUID
export const ER_UUIDViewOnShow: React.FC<ViewProps<types.ER_UUID>> = ER_StringViewOnShow
export const ER_UUIDViewOnList: React.FC<ViewProps<types.ER_UUID>> = ER_StringViewOnList
export const ER_UUIDViewOnForm: React.FC = ER_StringListViewOnForm
// endregion UUID

// region Code
export const ER_CodeViewOnShow: React.FC<ViewProps<types.ER_Code>> = ({value}) => {
    return <ER_CodeViewOnForm value={value} readOnly={true}/>
}
export const ER_CodeViewOnList: React.FC<ViewProps<types.ER_Code>> = ({value}) => {
    return <RA.TextField value={value} code={true} ellipsis={true} style={ {width: '100px'} }/>
}
export const ER_CodeViewOnForm: React.FC<any> = (props) => {
    return <CodeEditor
        {...props}
        language={props.Language || 'js'}
        padding={15}
        style={ {
            fontSize: 12,
            backgroundColor: "#000",
            fontFamily: 'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace',
        } }
    />
}
// endregion Code

// region Rich Text
export const ER_RichTextViewOnShow: React.FC<ViewProps<types.ER_RichText>> = ({value})=>{
    return <div dangerouslySetInnerHTML={ {__html: value || ""} }></div>
}
export const ER_RichTextViewOnList: React.FC<ViewProps<types.ER_RichText>> = ER_StringViewOnList
export const ER_RichTextViewOnForm: React.FC = (props) => <ReactQuill {...props} theme="snow"/>
// endregion Rich Text

// region URL
export const ER_URLViewOnShow: React.FC<ViewProps<types.ER_URL>> = ({value}) => {
    return <RA.Button type="primary" href={ value } target="_blank" icon={<RA.Icons.LinkOutlined/>}>{ value }</RA.Button>
}
export const ER_URLViewOnList: React.FC<ViewProps<types.ER_URL>> = ({value}) => {
    return <RA.Button type="primary" href={ value } target="_blank" icon={<RA.Icons.LinkOutlined/>}>Open</RA.Button>
}
export const ER_URLViewOnForm: React.FC = ER_StringViewOnForm
// endregion URL

// region Enums

export const ER_Enums_ProcessStatusViewOnShow = ER_StringViewOnShow
export const ER_Enums_ProcessStatusViewOnList = ER_StringViewOnList
export const ER_Enums_ProcessStatusViewOnForm: React.FC = (props)=>{
    return  <RA.Select
        {...props}
        options={[
            {
                value: "none",
                label: "None",
            },
            {
                value: "done",
                label: "Done",
            },
            {
                value: "enqueued",
                label: "Enqueued",
            },
            {
                value: "in_progress",
                label: "In Progress",
            },
            {
                value: "failed",
                label: "Failed",
            },
        ]}
    />
}
// endregion Enums