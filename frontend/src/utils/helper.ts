export type Goly = {
	id?: number;
	redirect?: string;
	url?: string;
	clicked?: number;
};

export type updateFieldsType = {
	redirect: string;
	url: string;
	id: number;
};

export type createFieldsType = {
	redirect: string;
	url: string;
};
