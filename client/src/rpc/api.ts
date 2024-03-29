// @generated by protobuf-ts 2.2.3-alpha.1 with parameter client_none,generate_dependencies
// @generated from protobuf file "api.proto" (package "lunabrain", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message lunabrain.Query
 */
export interface Query {
    /**
     * @generated from protobuf field: string query = 1;
     */
    query: string;
}
/**
 * @generated from protobuf message lunabrain.Results
 */
export interface Results {
    /**
     * @generated from protobuf field: repeated lunabrain.StoredContent storedContent = 1;
     */
    storedContent: StoredContent[];
}
/**
 * @generated from protobuf message lunabrain.StoredContent
 */
export interface StoredContent {
    /**
     * @generated from protobuf field: lunabrain.Content content = 1;
     */
    content?: Content;
    /**
     * @generated from protobuf field: lunabrain.NormalizedContent normalized = 2;
     */
    normalized?: NormalizedContent;
}
/**
 * @generated from protobuf message lunabrain.NormalizedContent
 */
export interface NormalizedContent {
    /**
     * @generated from protobuf field: string data = 1;
     */
    data: string;
}
/**
 * @generated from protobuf message lunabrain.File
 */
export interface File {
    /**
     * @generated from protobuf field: bytes data = 1;
     */
    data: Uint8Array;
}
/**
 * @generated from protobuf message lunabrain.ContentID
 */
export interface ContentID {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
}
/**
 * @generated from protobuf message lunabrain.Metadata
 */
export interface Metadata {
}
/**
 * Content has data and metadata
 *
 * @generated from protobuf message lunabrain.Content
 */
export interface Content {
    /**
     * @generated from protobuf field: bytes data = 1;
     */
    data: Uint8Array;
    /**
     * @generated from protobuf field: lunabrain.ContentType type = 2;
     */
    type: ContentType;
    /**
     * @generated from protobuf field: lunabrain.Metadata metadata = 3;
     */
    metadata?: Metadata;
    /**
     * @generated from protobuf field: string createdAt = 4;
     */
    createdAt: string;
}
/**
 * @generated from protobuf enum lunabrain.ContentType
 */
export enum ContentType {
    /**
     * @generated from protobuf enum value: TEXT = 0;
     */
    TEXT = 0,
    /**
     * @generated from protobuf enum value: AUDIO = 1;
     */
    AUDIO = 1,
    /**
     * @generated from protobuf enum value: URL = 2;
     */
    URL = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class Query$Type extends MessageType<Query> {
    constructor() {
        super("lunabrain.Query", [
            { no: 1, name: "query", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Query>): Query {
        const message = { query: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Query>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Query): Query {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string query */ 1:
                    message.query = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Query, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string query = 1; */
        if (message.query !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.query);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.Query
 */
export const Query = new Query$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Results$Type extends MessageType<Results> {
    constructor() {
        super("lunabrain.Results", [
            { no: 1, name: "storedContent", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => StoredContent }
        ]);
    }
    create(value?: PartialMessage<Results>): Results {
        const message = { storedContent: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Results>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Results): Results {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated lunabrain.StoredContent storedContent */ 1:
                    message.storedContent.push(StoredContent.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Results, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated lunabrain.StoredContent storedContent = 1; */
        for (let i = 0; i < message.storedContent.length; i++)
            StoredContent.internalBinaryWrite(message.storedContent[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.Results
 */
export const Results = new Results$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StoredContent$Type extends MessageType<StoredContent> {
    constructor() {
        super("lunabrain.StoredContent", [
            { no: 1, name: "content", kind: "message", T: () => Content },
            { no: 2, name: "normalized", kind: "message", T: () => NormalizedContent }
        ]);
    }
    create(value?: PartialMessage<StoredContent>): StoredContent {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<StoredContent>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StoredContent): StoredContent {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* lunabrain.Content content */ 1:
                    message.content = Content.internalBinaryRead(reader, reader.uint32(), options, message.content);
                    break;
                case /* lunabrain.NormalizedContent normalized */ 2:
                    message.normalized = NormalizedContent.internalBinaryRead(reader, reader.uint32(), options, message.normalized);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: StoredContent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* lunabrain.Content content = 1; */
        if (message.content)
            Content.internalBinaryWrite(message.content, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* lunabrain.NormalizedContent normalized = 2; */
        if (message.normalized)
            NormalizedContent.internalBinaryWrite(message.normalized, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.StoredContent
 */
export const StoredContent = new StoredContent$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NormalizedContent$Type extends MessageType<NormalizedContent> {
    constructor() {
        super("lunabrain.NormalizedContent", [
            { no: 1, name: "data", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<NormalizedContent>): NormalizedContent {
        const message = { data: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<NormalizedContent>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: NormalizedContent): NormalizedContent {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string data */ 1:
                    message.data = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: NormalizedContent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string data = 1; */
        if (message.data !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.data);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.NormalizedContent
 */
export const NormalizedContent = new NormalizedContent$Type();
// @generated message type with reflection information, may provide speed optimized methods
class File$Type extends MessageType<File> {
    constructor() {
        super("lunabrain.File", [
            { no: 1, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<File>): File {
        const message = { data: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<File>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: File): File {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes data */ 1:
                    message.data = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: File, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes data = 1; */
        if (message.data.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.data);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.File
 */
export const File = new File$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ContentID$Type extends MessageType<ContentID> {
    constructor() {
        super("lunabrain.ContentID", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ContentID>): ContentID {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ContentID>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ContentID): ContentID {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ContentID, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.ContentID
 */
export const ContentID = new ContentID$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Metadata$Type extends MessageType<Metadata> {
    constructor() {
        super("lunabrain.Metadata", []);
    }
    create(value?: PartialMessage<Metadata>): Metadata {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Metadata>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Metadata): Metadata {
        return target ?? this.create();
    }
    internalBinaryWrite(message: Metadata, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.Metadata
 */
export const Metadata = new Metadata$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Content$Type extends MessageType<Content> {
    constructor() {
        super("lunabrain.Content", [
            { no: 1, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "type", kind: "enum", T: () => ["lunabrain.ContentType", ContentType] },
            { no: 3, name: "metadata", kind: "message", T: () => Metadata },
            { no: 4, name: "createdAt", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Content>): Content {
        const message = { data: new Uint8Array(0), type: 0, createdAt: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Content>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Content): Content {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes data */ 1:
                    message.data = reader.bytes();
                    break;
                case /* lunabrain.ContentType type */ 2:
                    message.type = reader.int32();
                    break;
                case /* lunabrain.Metadata metadata */ 3:
                    message.metadata = Metadata.internalBinaryRead(reader, reader.uint32(), options, message.metadata);
                    break;
                case /* string createdAt */ 4:
                    message.createdAt = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Content, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes data = 1; */
        if (message.data.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.data);
        /* lunabrain.ContentType type = 2; */
        if (message.type !== 0)
            writer.tag(2, WireType.Varint).int32(message.type);
        /* lunabrain.Metadata metadata = 3; */
        if (message.metadata)
            Metadata.internalBinaryWrite(message.metadata, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* string createdAt = 4; */
        if (message.createdAt !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.createdAt);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message lunabrain.Content
 */
export const Content = new Content$Type();
/**
 * @generated ServiceType for protobuf service lunabrain.API
 */
export const API = new ServiceType("lunabrain.API", [
    { name: "Save", options: {}, I: Content, O: ContentID },
    { name: "Search", options: {}, I: Query, O: Results }
]);
