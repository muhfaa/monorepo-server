function validatePhone(phone) {
    const regex = /^(^\+62|62|^08)(\d{3,4}-?){2}\d{3,4}$/g;
  
    return phone.match(regex);
}

function validateName(name) {
    const regex = /^[a-zA-Z ]{2,30}$/;
  
    return name.match(regex);
}

function validateRole(role) {

    return role === "admin" || role === "user"
    
}

function validatePassword(password) {
    
    return password.lenght === 4

}

module.exports = {
    validatePhone,
    validateName,
    validateRole,
    validatePassword
}
