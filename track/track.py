import yaml

d={
"project":{
    "title":"project-tracker",
    "Start-Date":"25/11/20",
    "End-date":"5/12/20",
    "Languages":["Go","html","python"]
    }
}

f=open('C:/Users/Documents/Langauges/rest api/Go/proj/proj3/.idea/track.yaml','w')
f.write(yaml.dump(d))
f.close