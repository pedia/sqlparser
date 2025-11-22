use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use std::ptr;

use sqlparser::dialect::{GenericDialect, SQLiteDialect};
use sqlparser::parser::Parser;

#[unsafe(no_mangle)]
extern "C" fn parse(csql: *const c_char) -> *mut c_char {
    let cs = unsafe { CStr::from_ptr(csql) };
    let sql = cs.to_string_lossy();

    let dialect = SQLiteDialect {};
    let result = Parser::parse_sql(&dialect, &sql);
    match result {
        Ok(statements) => {
            let serialized = serde_json::to_string(&statements).unwrap();
            let cs = CString::new(serialized).expect("CString::new");
            cs.into_raw()
        }
        Err(_e) => ptr::null_mut(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn free_rust_string(ptr: *mut c_char) {
    unsafe {
        if ptr.is_null() {
            return;
        }

        drop(CString::from_raw(ptr));
    }
}
