export namespace backend {
	
	export class ExcelData {
	    filename: string;
	    header: string[];
	    details: string[][];
	
	    static createFrom(source: any = {}) {
	        return new ExcelData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.filename = source["filename"];
	        this.header = source["header"];
	        this.details = source["details"];
	    }
	}

}

