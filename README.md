# ent-cli

## MoSCoW

### Mo (Must have) 
- [x] quickstart with minikube and tekton pipeline crd
- [x] ent bundle with podman -> replace with binary version obtained with pkg
- [x] clients to manage all Entando CRs
- [ ] provide executables 
- [ ] migrates the commands
- [x] modularity: kubectl-style
- [ ] strict versioning
- [ ] uses ko for build container images -> To handle several platform sha

### S (Should have)
- [x] cluster native, interaction with the Entando CRs
- [ ] a command for entando backup
- [ ] set cpu/ram pod limits
- [ ] set http/https
- [x] set DB type(with entando-clients) 
- [ ] set DB internal or external
- [ ] set internal or external keycloak 

### Co (Could have)
- [x] profile management
- [x] extensibility of plugins in a kubectl-style

### W (Won't have)
- [ ] external dependencies

### How to extend the cli through the use of plugins
- rename your executable as `ent-cli-[commandName]`
- move your executable to one of the dir in your system's `PATH`
