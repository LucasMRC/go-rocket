export namespace frontend {
	
	export class ScreenSize {
	    width: number;
	    height: number;
	
	    static createFrom(source: any = {}) {
	        return new ScreenSize(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}
	export class Screen {
	    isCurrent: boolean;
	    isPrimary: boolean;
	    width: number;
	    height: number;
	    size: ScreenSize;
	    physicalSize: ScreenSize;
	
	    static createFrom(source: any = {}) {
	        return new Screen(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isCurrent = source["isCurrent"];
	        this.isPrimary = source["isPrimary"];
	        this.width = source["width"];
	        this.height = source["height"];
	        this.size = this.convertValues(source["size"], ScreenSize);
	        this.physicalSize = this.convertValues(source["physicalSize"], ScreenSize);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace main {
	
	export class AppResponse___github_com_wailsapp_wails_v2_internal_frontend_Screen_ {
	    status: number;
	    response: frontend.Screen[];
	
	    static createFrom(source: any = {}) {
	        return new AppResponse___github_com_wailsapp_wails_v2_internal_frontend_Screen_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.response = this.convertValues(source["response"], frontend.Screen);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class AppResponse_interface____ {
	    status: number;
	    response: any;
	
	    static createFrom(source: any = {}) {
	        return new AppResponse_interface____(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.response = source["response"];
	    }
	}

}

