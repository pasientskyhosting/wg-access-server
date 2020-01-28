// Generated by protoc-gen-grpc-ts-web. DO NOT EDIT!
/* eslint-disable */
/* tslint:disable */

import * as jspb from 'google-protobuf';
import * as grpcWeb from 'grpc-web';

import * as googleProtobufWrappers from 'google-protobuf/google/protobuf/wrappers_pb';

export class Server {

	private client_ = new grpcWeb.GrpcWebClientBase({
		format: 'text',
	});

	private methodInfoInfo = new grpcWeb.AbstractClientBase.MethodInfo(
		InfoRes,
		(req: InfoReq) => req.serializeBinary(),
		InfoRes.deserializeBinary
	);

	constructor(
		private hostname: string,
		private defaultMetadata?: () => grpcWeb.Metadata,
	) { }

	info(req: InfoReq.AsObject, metadata?: grpcWeb.Metadata): Promise<InfoRes.AsObject> {
		return new Promise((resolve, reject) => {
			const message = InfoReqFromObject(req);
			this.client_.rpcCall(
				this.hostname + '/proto.Server/Info',
				message,
				Object.assign({}, this.defaultMetadata ? this.defaultMetadata() : {}, metadata),
				this.methodInfoInfo,
				(err: grpcWeb.Error, res: InfoRes) => {
					if (err) {
						reject(err);
					} else {
						resolve(res.toObject());
					}
				},
			);
		});
	}

}




export declare namespace InfoReq {
	export type AsObject = {
	}
}

export class InfoReq extends jspb.Message {

	private static repeatedFields_ = [
		
	];

	constructor(data?: jspb.Message.MessageArray) {
		super();
		jspb.Message.initialize(this, data || [], 0, -1, InfoReq.repeatedFields_, null);
	}


	serializeBinary(): Uint8Array {
		const writer = new jspb.BinaryWriter();
		InfoReq.serializeBinaryToWriter(this, writer);
		return writer.getResultBuffer();
	}

	toObject(): InfoReq.AsObject {
		let f: any;
		return {
		};
	}

	static serializeBinaryToWriter(message: InfoReq, writer: jspb.BinaryWriter): void {
	}

	static deserializeBinary(bytes: Uint8Array): InfoReq {
		var reader = new jspb.BinaryReader(bytes);
		var message = new InfoReq();
		return InfoReq.deserializeBinaryFromReader(message, reader);
	}

	static deserializeBinaryFromReader(message: InfoReq, reader: jspb.BinaryReader): InfoReq {
		while (reader.nextField()) {
			if (reader.isEndGroup()) {
				break;
			}
			const field = reader.getFieldNumber();
			switch (field) {
			default:
				reader.skipField();
				break;
			}
		}
		return message;
	}

}
export declare namespace InfoRes {
	export type AsObject = {
		publicKey: string,
		host?: googleProtobufWrappers.StringValue.AsObject,
		port: number,
		hostVpnIp: string,
	}
}

export class InfoRes extends jspb.Message {

	private static repeatedFields_ = [
		
	];

	constructor(data?: jspb.Message.MessageArray) {
		super();
		jspb.Message.initialize(this, data || [], 0, -1, InfoRes.repeatedFields_, null);
	}


	getPublicKey(): string {
		return jspb.Message.getFieldWithDefault(this, 1, "");
	}

	setPublicKey(value: string): void {
		(jspb.Message as any).setProto3StringField(this, 1, value);
	}

	getHost(): googleProtobufWrappers.StringValue {
		return jspb.Message.getWrapperField(this, googleProtobufWrappers.StringValue, 2);
	}

	setHost(value?: googleProtobufWrappers.StringValue): void {
		(jspb.Message as any).setWrapperField(this, 2, value);
	}

	getPort(): number {
		return jspb.Message.getFieldWithDefault(this, 3, 0);
	}

	setPort(value: number): void {
		(jspb.Message as any).setProto3IntField(this, 3, value);
	}

	getHostVpnIp(): string {
		return jspb.Message.getFieldWithDefault(this, 4, "");
	}

	setHostVpnIp(value: string): void {
		(jspb.Message as any).setProto3StringField(this, 4, value);
	}

	serializeBinary(): Uint8Array {
		const writer = new jspb.BinaryWriter();
		InfoRes.serializeBinaryToWriter(this, writer);
		return writer.getResultBuffer();
	}

	toObject(): InfoRes.AsObject {
		let f: any;
		return {publicKey: this.getPublicKey(),
			host: (f = this.getHost()) && f.toObject(),
			port: this.getPort(),
			hostVpnIp: this.getHostVpnIp(),
			
		};
	}

	static serializeBinaryToWriter(message: InfoRes, writer: jspb.BinaryWriter): void {
		const field1 = message.getPublicKey();
		if (field1.length > 0) {
			writer.writeString(1, field1);
		}
		const field2 = message.getHost();
		if (field2 != null) {
			writer.writeMessage(2, field2, googleProtobufWrappers.StringValue.serializeBinaryToWriter);
		}
		const field3 = message.getPort();
		if (field3 != 0) {
			writer.writeInt32(3, field3);
		}
		const field4 = message.getHostVpnIp();
		if (field4.length > 0) {
			writer.writeString(4, field4);
		}
	}

	static deserializeBinary(bytes: Uint8Array): InfoRes {
		var reader = new jspb.BinaryReader(bytes);
		var message = new InfoRes();
		return InfoRes.deserializeBinaryFromReader(message, reader);
	}

	static deserializeBinaryFromReader(message: InfoRes, reader: jspb.BinaryReader): InfoRes {
		while (reader.nextField()) {
			if (reader.isEndGroup()) {
				break;
			}
			const field = reader.getFieldNumber();
			switch (field) {
			case 1:
				const field1 = reader.readString()
				message.setPublicKey(field1);
				break;
			case 2:
				const field2 = new googleProtobufWrappers.StringValue();
				reader.readMessage(field2, googleProtobufWrappers.StringValue.deserializeBinaryFromReader);
				message.setHost(field2);
				break;
			case 3:
				const field3 = reader.readInt32()
				message.setPort(field3);
				break;
			case 4:
				const field4 = reader.readString()
				message.setHostVpnIp(field4);
				break;
			default:
				reader.skipField();
				break;
			}
		}
		return message;
	}

}


function InfoReqFromObject(obj: InfoReq.AsObject | undefined): InfoReq | undefined {
	if (obj === undefined) {
		return undefined;
	}
	const message = new InfoReq();
	return message;
}

function InfoResFromObject(obj: InfoRes.AsObject | undefined): InfoRes | undefined {
	if (obj === undefined) {
		return undefined;
	}
	const message = new InfoRes();
	message.setPublicKey(obj.publicKey);
	message.setHost(StringValueFromObject(obj.host));
	message.setPort(obj.port);
	message.setHostVpnIp(obj.hostVpnIp);
	return message;
}

function StringValueFromObject(obj: googleProtobufWrappers.StringValue.AsObject | undefined): googleProtobufWrappers.StringValue | undefined {
	if (obj === undefined) {
		return undefined;
	}
	const message = new googleProtobufWrappers.StringValue();
	message.setValue(obj.value);
	return message;
}

