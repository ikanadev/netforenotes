/* tslint:disable */
/* eslint-disable */
/**
 * Netforenotes app API docs
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface NoteBase
 */
export interface NoteBase {
    /**
     * 
     * @type {string}
     * @memberof NoteBase
     */
    title: string;
    /**
     * 
     * @type {number}
     * @memberof NoteBase
     */
    createdAt: number;
}

/**
 * Check if a given object implements the NoteBase interface.
 */
export function instanceOfNoteBase(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "title" in value;
    isInstance = isInstance && "createdAt" in value;

    return isInstance;
}

export function NoteBaseFromJSON(json: any): NoteBase {
    return NoteBaseFromJSONTyped(json, false);
}

export function NoteBaseFromJSONTyped(json: any, ignoreDiscriminator: boolean): NoteBase {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'title': json['title'],
        'createdAt': json['created_at'],
    };
}

export function NoteBaseToJSON(value?: NoteBase | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'title': value.title,
        'created_at': value.createdAt,
    };
}

