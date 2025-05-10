export namespace backend {
	
	export class ExcelData {
	    filename: string;
	    isSaveDB: boolean;
	    header: string[];
	    details: string[][];
	
	    static createFrom(source: any = {}) {
	        return new ExcelData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.filename = source["filename"];
	        this.isSaveDB = source["isSaveDB"];
	        this.header = source["header"];
	        this.details = source["details"];
	    }
	}

}

export namespace models {
	
	export class Response {
	    status: string;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}

}

