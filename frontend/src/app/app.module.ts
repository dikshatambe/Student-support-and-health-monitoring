import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { HomeComponent } from './home/home.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { MatInputModule } from '@angular/material/input';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatIconModule } from '@angular/material/icon'
import { MatCardModule } from '@angular/material/card'
import { MatButtonModule } from '@angular/material/button'
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner'
import { MatGridListModule } from '@angular/material/grid-list';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSnackBarModule } from "@angular/material/snack-bar";
import { RegisterUserComponent } from './register-user/register-user.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
//import { UpdateuserComponent } from './updateuser/updateuser.component';
import { SearchComponent } from './search/search.component';
import { DeleteComponent } from './delete/delete.component';
import { UpdateComponent } from './update/update.component';
import { MatSidenavModule } from '@angular/material/sidenav';
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



@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    RegisterUserComponent,
    SearchComponent,
    DeleteComponent,
    UpdateComponent,
    LoginComponent,
    GroupComponent,
    AdminComponent,
    UserComponent,

    CreateuserComponent,
    DeleteuserComponent,
    SearchuserComponent,
    UpdateuserComponent,
    HelpComponent,
    HealthComponent,
    UserhealthComponent


  ],
  imports: [
    MatSidenavModule,
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatCardModule,
    MatProgressSpinnerModule,
    MatGridListModule,
    MatDialogModule,
    MatFormFieldModule,
    MatSnackBarModule,
    MatInputModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
