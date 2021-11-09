function login(){
    const form = document.login_form;
    const chkUsername = checkValidUsername(form);
    const chkPw = checkVlidPassword(form);

    if (chkUsername && chkPw){
        console.log('complete. form.submit();');
    }
}

function getValue(form){

}
function checkValidUsername(form){
    if(form.username.value == ""){
        document.getElementById('alert_username').innerText="enter username";
        return false;
    }
    return true;
}
function checkValidPassword(form){
    if(form.password.value == ""){
        document.getElementById('alert_password').innerText="enter password";
        return false;
    }
    return true;
}
function check_input(form){
    if (!document.login_form.id_val.value)
    {
        alert("아이디를 다시 입력 하세요");
        document.login_form.id_val.focus();
        return;
    }
    if (!document.login_form.pw_val.value)
    {
        alert("비밀번호를 다시 입력하세요.");
        return;
    }
    if (document.login_form.pw_val.value == "abc" && document.login_form.id_val.value == "123"){
        alert("로그인 되었습니다.");
        document.login_form.submit();
    }
    else{
        alert("ID나 PW가 틀렸습니다.")
    }
}