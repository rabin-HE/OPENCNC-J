export default ({
    status: (onload: any, onerror: any) =>{
        let image = new Image();
        image.onload = function(){
            if(typeof onload == "function"){
                onload();
            }
        };
        image.onerror = function(){
            if(typeof onerror == "function"){
                onerror();
            }
        };
        image.src = "https://cdn.geekros.com/studio/network/network.png";
    }
});