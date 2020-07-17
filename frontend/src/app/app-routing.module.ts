import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { SearchComponent } from './search/search.component';
import { DeleteComponent } from './delete/delete.component';
import { UpdateComponent } from './update/update.component';
import { RegisterUserComponent } from './register-user/register-user.component';
import { LoginComponent } from './login/login.component';
import { GroupComponent } from './group/group.component';
import { AdminComponent } from './admin/admin.component';
import { UserComponent } from './user/user.component';

import { CreateuserComponent } from './UserOperations/createuser/createuser.component';
import { DeleteuserComponent } from './UserOperations/deleteuser/deleteuser.component';
import { SearchuserComponent } from './UserOperations/searchuser/searchuser.component';
import { UpdateuserComponent } from './UserOperations/updateuser/updateuser.component';
import { HelpComponent } from './help/help.component';
import { HealthComponent } from './health/health.component';
import { UserhealthComponent } from './userhealth/userhealth.component';

const routes: Routes = [{ path: '', redirectTo: 'login', pathMatch: 'full' },
{
  path: 'home', component: HomeComponent,
  // canActivate: [AuthenticationGuard],
},
{ path: 'login', component: LoginComponent, runGuardsAndResolvers: 'always', },
{ path: 'group', component: GroupComponent },
{ path: 'admin', component: AdminComponent },
{ path: 'user', component: UserComponent },

{ path: 'createuser', component: CreateuserComponent },
{ path: 'deleteuser', component: DeleteuserComponent },
{ path: 'searchuser', component: SearchuserComponent },
{ path: 'register', component: RegisterUserComponent },
{ path: 'updateuser', component: UpdateuserComponent },
{ path: 'search', component: SearchComponent },
{ path: 'delete', component: DeleteComponent },
{ path: 'update', component: UpdateComponent },
{ path: 'help', component: HelpComponent },
{ path: 'health', component: HealthComponent },
{ path: 'userhealth', component: UserhealthComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { onSameUrlNavigation: "reload" })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
