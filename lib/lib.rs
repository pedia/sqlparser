use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use std::ptr;

use sqlparser::dialect::dialect_from_str;
use sqlparser::parser::Parser;

#[unsafe(no_mangle)]
pub extern "C" fn parse(cdialect: *const c_char, csql: *const c_char) -> *mut c_char {
    let dialect_name = unsafe { CStr::from_ptr(cdialect) }.to_string_lossy();
    let dialect = dialect_from_str(dialect_name);
    if dialect.is_none() {
        return ptr::null_mut()
    }

    let dialect = dialect.unwrap();
    let sql = unsafe { CStr::from_ptr(csql) }.to_string_lossy();
    let result = Parser::parse_sql(&*dialect, &sql);
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
