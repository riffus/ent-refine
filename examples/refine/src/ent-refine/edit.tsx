import {useState} from "react";
import * as RA from "@pankod/refine-antd";
import * as Interfaces from "./interfaces";
import {Cursors} from "./data-provider";
import dayjs from "dayjs";
import * as FieldView from "./field-view";
import * as Custom from "./custom";

import CodeEditor from '@uiw/react-textarea-code-editor';

import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';

export const CompanyEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.ICompany>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "description",
                    {
                        "product": [ "id" ]
                    },
                    {
                        "countries": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "phones": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "emails": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "websites": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "locations": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "logoImage": [ "id" ]
                    },
                    {
                        "galleryImages": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                ],
            }
        }
    );
    const [ productCursors, setProductCursors] = useState<Cursors>({})
    const { selectProps: productSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Product",
        optionLabel: "url",
        optionValue: "id",
        metaData:{
            cursors: productCursors,
            fields: ["id", "url"]
        },
        onSearch: (value) => [
            {
                field: "url",
                operator: "contains",
                value,
            },
        ],
    });
    const [ countriesCursors, setCountriesCursors] = useState<Cursors>({})
    const { selectProps: countriesSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Country",
        optionLabel: "name",
        optionValue: "id",
        metaData:{
            cursors: countriesCursors,
            fields: ["id", "name"]
        },
        onSearch: (value) => [
            {
                field: "name",
                operator: "contains",
                value,
            },
        ],
    });
    const [ phonesCursors, setPhonesCursors] = useState<Cursors>({})
    const { selectProps: phonesSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Phone",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: phonesCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ emailsCursors, setEmailsCursors] = useState<Cursors>({})
    const { selectProps: emailsSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Email",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: emailsCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ websitesCursors, setWebsitesCursors] = useState<Cursors>({})
    const { selectProps: websitesSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Website",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: websitesCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ locationsCursors, setLocationsCursors] = useState<Cursors>({})
    const { selectProps: locationsSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Location",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: locationsCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ logo_imageCursors, setLogoImageCursors] = useState<Cursors>({})
    const { selectProps: logo_imageSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Image",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: logo_imageCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ gallery_imagesCursors, setGalleryImagesCursors] = useState<Cursors>({})
    const { selectProps: gallery_imagesSelectProps } = RA.useSelect<Interfaces.ICompany>({
        resource: "Image",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: gallery_imagesCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Description"
                    name="description"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_RichTextViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="product" name="productID" rules={[{required: false}]}>
                    <RA.Select {...productSelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="countries" name={["countryIDs"]} rules={[{required: false}]}>
                    <RA.Select {...countriesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="phones" name={["phoneIDs"]} rules={[{required: false}]}>
                    <RA.Select {...phonesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="emails" name={["emailIDs"]} rules={[{required: false}]}>
                    <RA.Select {...emailsSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="websites" name={["websiteIDs"]} rules={[{required: false}]}>
                    <RA.Select {...websitesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="locations" name={["locationIDs"]} rules={[{required: false}]}>
                    <RA.Select {...locationsSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="logo_image" name="logoImageID" rules={[{required: false}]}>
                    <RA.Select {...logo_imageSelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="gallery_images" name={["galleryImageIDs"]} rules={[{required: false}]}>
                    <RA.Select {...gallery_imagesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const CountryEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.ICountry>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "name",
                    "code",
                    {
                        "companies": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "phones": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "emails": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "websites": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "locations": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                ],
            }
        }
    );
    const [ companiesCursors, setCompaniesCursors] = useState<Cursors>({})
    const { selectProps: companiesSelectProps } = RA.useSelect<Interfaces.ICountry>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: companiesCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ phonesCursors, setPhonesCursors] = useState<Cursors>({})
    const { selectProps: phonesSelectProps } = RA.useSelect<Interfaces.ICountry>({
        resource: "Phone",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: phonesCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ emailsCursors, setEmailsCursors] = useState<Cursors>({})
    const { selectProps: emailsSelectProps } = RA.useSelect<Interfaces.ICountry>({
        resource: "Email",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: emailsCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ websitesCursors, setWebsitesCursors] = useState<Cursors>({})
    const { selectProps: websitesSelectProps } = RA.useSelect<Interfaces.ICountry>({
        resource: "Website",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: websitesCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ locationsCursors, setLocationsCursors] = useState<Cursors>({})
    const { selectProps: locationsSelectProps } = RA.useSelect<Interfaces.ICountry>({
        resource: "Location",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: locationsCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Name"
                    name="name"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Code"
                    name="code"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="companies" name={["companyIDs"]} rules={[{required: false}]}>
                    <RA.Select {...companiesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="phones" name={["phoneIDs"]} rules={[{required: false}]}>
                    <RA.Select {...phonesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="emails" name={["emailIDs"]} rules={[{required: false}]}>
                    <RA.Select {...emailsSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="websites" name={["websiteIDs"]} rules={[{required: false}]}>
                    <RA.Select {...websitesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="locations" name={["locationIDs"]} rules={[{required: false}]}>
                    <RA.Select {...locationsSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const EmailEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.IEmail>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "description",
                    "address",
                    {
                        "company": [ "id" ]
                    },
                    {
                        "country": [ "id" ]
                    },
                ],
            }
        }
    );
    const [ companyCursors, setCompanyCursors] = useState<Cursors>({})
    const { selectProps: companySelectProps } = RA.useSelect<Interfaces.IEmail>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: companyCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ countryCursors, setCountryCursors] = useState<Cursors>({})
    const { selectProps: countrySelectProps } = RA.useSelect<Interfaces.IEmail>({
        resource: "Country",
        optionLabel: "name",
        optionValue: "id",
        metaData:{
            cursors: countryCursors,
            fields: ["id", "name"]
        },
        onSearch: (value) => [
            {
                field: "name",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Description"
                    name="description"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Address"
                    name="address"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="company" name="companyID" rules={[{required: false}]}>
                    <RA.Select {...companySelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="country" name="countryID" rules={[{required: false}]}>
                    <RA.Select {...countrySelectProps} mode={ undefined }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const ImageEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.IImage>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "originalURL",
                    {
                        "galleryCompany": [ "id" ]
                    },
                    {
                        "logoCompany": [ "id" ]
                    },
                ],
            }
        }
    );
    const [ gallery_companyCursors, setGalleryCompanyCursors] = useState<Cursors>({})
    const { selectProps: gallery_companySelectProps } = RA.useSelect<Interfaces.IImage>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: gallery_companyCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ logo_companyCursors, setLogoCompanyCursors] = useState<Cursors>({})
    const { selectProps: logo_companySelectProps } = RA.useSelect<Interfaces.IImage>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: logo_companyCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Original Url"
                    name="originalURL"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_ImageViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="gallery_company" name="galleryCompanyID" rules={[{required: false}]}>
                    <RA.Select {...gallery_companySelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="logo_company" name="logoCompanyID" rules={[{required: false}]}>
                    <RA.Select {...logo_companySelectProps} mode={ undefined }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const LocationEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.ILocation>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "description",
                    "latitude",
                    "longitude",
                    "address",
                    "postcode",
                    "type",
                    "state",
                    "suburb",
                    "streetType",
                    "streetName",
                    {
                        "company": [ "id" ]
                    },
                    {
                        "country": [ "id" ]
                    },
                ],
            }
        }
    );
    const [ companyCursors, setCompanyCursors] = useState<Cursors>({})
    const { selectProps: companySelectProps } = RA.useSelect<Interfaces.ILocation>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: companyCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ countryCursors, setCountryCursors] = useState<Cursors>({})
    const { selectProps: countrySelectProps } = RA.useSelect<Interfaces.ILocation>({
        resource: "Country",
        optionLabel: "name",
        optionValue: "id",
        metaData:{
            cursors: countryCursors,
            fields: ["id", "name"]
        },
        onSearch: (value) => [
            {
                field: "name",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Description"
                    name="description"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Latitude"
                    name="latitude"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Longitude"
                    name="longitude"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Address"
                    name="address"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Postcode"
                    name="postcode"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Type"
                    name="type"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="State"
                    name="state"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Suburb"
                    name="suburb"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Street Type"
                    name="streetType"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Street Name"
                    name="streetName"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="company" name="companyID" rules={[{required: false}]}>
                    <RA.Select {...companySelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="country" name="countryID" rules={[{required: false}]}>
                    <RA.Select {...countrySelectProps} mode={ undefined }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const PhoneEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.IPhone>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "description",
                    "number",
                    "type",
                    {
                        "company": [ "id" ]
                    },
                    {
                        "country": [ "id" ]
                    },
                ],
            }
        }
    );
    const [ companyCursors, setCompanyCursors] = useState<Cursors>({})
    const { selectProps: companySelectProps } = RA.useSelect<Interfaces.IPhone>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: companyCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ countryCursors, setCountryCursors] = useState<Cursors>({})
    const { selectProps: countrySelectProps } = RA.useSelect<Interfaces.IPhone>({
        resource: "Country",
        optionLabel: "name",
        optionValue: "id",
        metaData:{
            cursors: countryCursors,
            fields: ["id", "name"]
        },
        onSearch: (value) => [
            {
                field: "name",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Description"
                    name="description"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Number"
                    name="number"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Type"
                    name="type"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="company" name="companyID" rules={[{required: false}]}>
                    <RA.Select {...companySelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="country" name="countryID" rules={[{required: false}]}>
                    <RA.Select {...countrySelectProps} mode={ undefined }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const ProductEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.IProduct>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "description",
                    "image",
                    "url",
                    "lastSell",
                    "createdAt",
                    "status",
                    "buildStatus",
                    {
                        "company": [ "id" ]
                    },
                    {
                        "warehouse": [ "id" ]
                    },
                    {
                        "vendor": [ "id" ]
                    },
                ],
            }
        }
    );
    const [ companyCursors, setCompanyCursors] = useState<Cursors>({})
    const { selectProps: companySelectProps } = RA.useSelect<Interfaces.IProduct>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: companyCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ warehouseCursors, setWarehouseCursors] = useState<Cursors>({})
    const { selectProps: warehouseSelectProps } = RA.useSelect<Interfaces.IProduct>({
        resource: "Warehouse",
        optionLabel: "url",
        optionValue: "id",
        metaData:{
            cursors: warehouseCursors,
            fields: ["id", "url"]
        },
        onSearch: (value) => [
            {
                field: "url",
                operator: "contains",
                value,
            },
        ],
    });
    const [ vendorCursors, setVendorCursors] = useState<Cursors>({})
    const { selectProps: vendorSelectProps } = RA.useSelect<Interfaces.IProduct>({
        resource: "Vendor",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: vendorCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Description"
                    name="description"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_RichTextViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Image"
                    name="image"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_ImageViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Url"
                    name="url"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_URLViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Last Sell"
                    name="lastSell"
                    rules={[{required: false}]}
                    getValueProps={(value) => ({
                        value: value ? dayjs(value) : "",
                    })}
                >
                    <FieldView.ER_DateViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Created At"
                    name="createdAt"
                    rules={[{required: false}]}
                    getValueProps={(value) => ({
                        value: value ? dayjs(value) : "",
                    })}
                >
                    <FieldView.ER_DateViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Status"
                    name="status"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_Enums_ProcessStatusViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Build Status"
                    name="buildStatus"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_Enums_ProcessStatusViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="company" name="companyID" rules={[{required: false}]}>
                    <RA.Select {...companySelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="warehouse" name="warehouseID" rules={[{required: false}]}>
                    <RA.Select {...warehouseSelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="vendor" name="vendorID" rules={[{required: false}]}>
                    <RA.Select {...vendorSelectProps} mode={ undefined }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const VendorEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.IVendor>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "url",
                    "schema",
                    {
                        "warehouses": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "products": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                ],
            }
        }
    );
    const [ warehousesCursors, setWarehousesCursors] = useState<Cursors>({})
    const { selectProps: warehousesSelectProps } = RA.useSelect<Interfaces.IVendor>({
        resource: "Warehouse",
        optionLabel: "url",
        optionValue: "id",
        metaData:{
            cursors: warehousesCursors,
            fields: ["id", "url"]
        },
        onSearch: (value) => [
            {
                field: "url",
                operator: "contains",
                value,
            },
        ],
    });
    const [ productsCursors, setProductsCursors] = useState<Cursors>({})
    const { selectProps: productsSelectProps } = RA.useSelect<Interfaces.IVendor>({
        resource: "Product",
        optionLabel: "url",
        optionValue: "id",
        metaData:{
            cursors: productsCursors,
            fields: ["id", "url"]
        },
        onSearch: (value) => [
            {
                field: "url",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Url"
                    name="url"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Schema"
                    name="schema"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_CodeViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="warehouses" name={["warehouseIDs"]} rules={[{required: false}]}>
                    <RA.Select {...warehousesSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="products" name={["productIDs"]} rules={[{required: false}]}>
                    <RA.Select {...productsSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const WarehouseEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.IWarehouse>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "url",
                    "lastUpdate",
                    "originalData",
                    "enabled",
                    "filters",
                    {
                        "products": [
                            {
                                edges: [
                                    {
                                        node: [ "id" ],
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "vendor": [ "id" ]
                    },
                ],
            }
        }
    );
    const [ productsCursors, setProductsCursors] = useState<Cursors>({})
    const { selectProps: productsSelectProps } = RA.useSelect<Interfaces.IWarehouse>({
        resource: "Product",
        optionLabel: "url",
        optionValue: "id",
        metaData:{
            cursors: productsCursors,
            fields: ["id", "url"]
        },
        onSearch: (value) => [
            {
                field: "url",
                operator: "contains",
                value,
            },
        ],
    });
    const [ vendorCursors, setVendorCursors] = useState<Cursors>({})
    const { selectProps: vendorSelectProps } = RA.useSelect<Interfaces.IWarehouse>({
        resource: "Vendor",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: vendorCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Url"
                    name="url"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Last Update"
                    name="lastUpdate"
                    rules={[{required: false}]}
                    getValueProps={(value) => ({
                        value: value ? dayjs(value) : "",
                    })}
                >
                    <FieldView.ER_DateViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Original Data"
                    name="originalData"
                    rules={[{required: false}]}
                >
                    <FieldView.ER_CodeViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Enabled"
                    name="enabled"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_BooleanViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Filters"
                    name="filters"
                    rules={[{required: false}]}
                >
                    <FieldView.ER_StringListViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="products" name={["productIDs"]} rules={[{required: false}]}>
                    <RA.Select {...productsSelectProps} mode={ "multiple" }/>
                </RA.Form.Item>
                <RA.Form.Item label="vendor" name="vendorID" rules={[{required: false}]}>
                    <RA.Select {...vendorSelectProps} mode={ undefined }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};

export const WebsiteEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } = RA.useForm<Interfaces.IWebsite>(
        {
            redirect: false,
            metaData: {
                fields: [
                    "title",
                    "description",
                    "url",
                    {
                        "company": [ "id" ]
                    },
                    {
                        "country": [ "id" ]
                    },
                ],
            }
        }
    );
    const [ companyCursors, setCompanyCursors] = useState<Cursors>({})
    const { selectProps: companySelectProps } = RA.useSelect<Interfaces.IWebsite>({
        resource: "Company",
        optionLabel: "title",
        optionValue: "id",
        metaData:{
            cursors: companyCursors,
            fields: ["id", "title"]
        },
        onSearch: (value) => [
            {
                field: "title",
                operator: "contains",
                value,
            },
        ],
    });
    const [ countryCursors, setCountryCursors] = useState<Cursors>({})
    const { selectProps: countrySelectProps } = RA.useSelect<Interfaces.IWebsite>({
        resource: "Country",
        optionLabel: "name",
        optionValue: "id",
        metaData:{
            cursors: countryCursors,
            fields: ["id", "name"]
        },
        onSearch: (value) => [
            {
                field: "name",
                operator: "contains",
                value,
            },
        ],
    });

    return (
        <RA.Edit saveButtonProps={saveButtonProps}>
            <RA.Form {...formProps} layout="vertical">
                
                <RA.Form.Item
                    label="Title"
                    name="title"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Description"
                    name="description"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                
                <RA.Form.Item
                    label="Url"
                    name="url"
                    rules={[{required: true}]}
                >
                    <FieldView.ER_StringViewOnForm/>
                </RA.Form.Item>
                <RA.Form.Item label="company" name="companyID" rules={[{required: false}]}>
                    <RA.Select {...companySelectProps} mode={ undefined }/>
                </RA.Form.Item>
                <RA.Form.Item label="country" name="countryID" rules={[{required: false}]}>
                    <RA.Select {...countrySelectProps} mode={ undefined }/>
                </RA.Form.Item>
            </RA.Form>
        </RA.Edit>
    );
};