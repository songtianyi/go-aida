/**
 * Created by Luna_Shu on 2017/7/28.
 *
 * @author Luna_Shu
 *
 */

// var Login = {template: '<div}}</div>'};
var Manage = {template: '<div>manage</div>'};

var routes = [
    {
        path: '/',
        component: Login
    },
    {
        path: '/manage',
        component: Manage
    }
];

router = new VueRouter({
    routes: routes
});
