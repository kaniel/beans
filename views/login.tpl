{{template "base/base.html" .}}
{{define "head"}}
<title> test-login </title>
<script type="text/javascript">
$(function(){
    $("#ddd").dialog({
        closable:false,
        buttons:[{
        text:'登录111',
        iconCls:'icon-save',
        handler:function(){
            fromsubmit();
        }
    },{
        text:'重置',
        iconCls:'icon-save',
        handler:function(){
            $("#form").from("reset");
        }
    }]
    });
});


function fromsubmit(){
    $("#form").form('submit',{
        url:'/auth/login',
        onSubmit:function(){
            return $("#form").form('validate');
        },
        success:function(r){
            location.href = "/"
            // var r = $.parseJSON( r );
            //if(r.status){
            //    location.href = "/auth"
            //}else{
            //    alert(r.info);
            //}
        }
    });
}
    //这个就是键盘触发的函数
var SubmitOrHidden = function(evt){
    evt = window.event || evt;
    if(evt.keyCode==13){//如果取到的键值是回车
          fromsubmit();
     }

}
</script>
{{end}}
{{define "body"}}
    {{define ""}}
<div style="text-align:center;margin:0 auto;width:350px;height:250px;" id="ddd" title="登录">
    <div style="padding:20px 20px 20px 40px;" >
        <form id="form" method="post">
            <table >
                <tr>
                    <td>用户名：</td><td><input type="text" class="easyui-validatebox" required="true" name="username" missingMessage="请输入用户名"/></td>
                </tr>
                <tr>
                    <td>密码：</td><td><input type="password" class="easyui-validatebox" required="true" name="password" missingMessage="请输入密码"/></td>
                </tr>
            </table>
        </form>
    </div>
</div>
{{end}}
