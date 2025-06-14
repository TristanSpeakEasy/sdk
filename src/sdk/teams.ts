/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { teamsCreateTeam } from "../funcs/teamsCreateTeam.js";
import { teamsDeleteTeam } from "../funcs/teamsDeleteTeam.js";
import { teamsDeleteTeamInviteCode } from "../funcs/teamsDeleteTeamInviteCode.js";
import { teamsGetTeam } from "../funcs/teamsGetTeam.js";
import { teamsGetTeamAccessRequest } from "../funcs/teamsGetTeamAccessRequest.js";
import { teamsGetTeamMembers } from "../funcs/teamsGetTeamMembers.js";
import { teamsGetTeams } from "../funcs/teamsGetTeams.js";
import { teamsInviteUserToTeam } from "../funcs/teamsInviteUserToTeam.js";
import { teamsJoinTeam } from "../funcs/teamsJoinTeam.js";
import { teamsPatchTeam } from "../funcs/teamsPatchTeam.js";
import { teamsRemoveTeamMember } from "../funcs/teamsRemoveTeamMember.js";
import { teamsRequestAccessToTeam } from "../funcs/teamsRequestAccessToTeam.js";
import { teamsUpdateTeamMember } from "../funcs/teamsUpdateTeamMember.js";
import { ClientSDK, RequestOptions } from "../lib/sdks.js";
import {
  CreateTeamRequestBody,
  CreateTeamResponseBody,
} from "../models/createteamop.js";
import {
  DeleteTeamInviteCodeRequest,
  DeleteTeamInviteCodeResponseBody,
} from "../models/deleteteaminvitecodeop.js";
import {
  DeleteTeamRequest,
  DeleteTeamResponseBody,
} from "../models/deleteteamop.js";
import {
  GetTeamAccessRequestRequest,
  GetTeamAccessRequestResponseBody,
} from "../models/getteamaccessrequestop.js";
import {
  GetTeamMembersRequest,
  GetTeamMembersResponseBody,
} from "../models/getteammembersop.js";
import { GetTeamRequest } from "../models/getteamop.js";
import { GetTeamsRequest, GetTeamsResponseBody } from "../models/getteamsop.js";
import {
  InviteUserToTeamRequest,
  InviteUserToTeamResponseBody,
} from "../models/inviteusertoteamop.js";
import { JoinTeamRequest, JoinTeamResponseBody } from "../models/jointeamop.js";
import { PatchTeamRequest } from "../models/patchteamop.js";
import {
  RemoveTeamMemberRequest,
  RemoveTeamMemberResponseBody,
} from "../models/removeteammemberop.js";
import {
  RequestAccessToTeamRequest,
  RequestAccessToTeamResponseBody,
} from "../models/requestaccesstoteamop.js";
import { Team } from "../models/team.js";
import {
  UpdateTeamMemberRequest,
  UpdateTeamMemberResponseBody,
} from "../models/updateteammemberop.js";
import { unwrapAsync } from "../types/fp.js";

export class Teams extends ClientSDK {
  /**
   * List team members
   *
   * @remarks
   * Get a paginated list of team members for the provided team.
   */
  async getTeamMembers(
    request: GetTeamMembersRequest,
    options?: RequestOptions,
  ): Promise<GetTeamMembersResponseBody> {
    return unwrapAsync(teamsGetTeamMembers(
      this,
      request,
      options,
    ));
  }

  /**
   * Invite a user
   *
   * @remarks
   * Invite a user to join the team specified in the URL. The authenticated user needs to be an `OWNER` in order to successfully invoke this endpoint. The user can be specified with an email or an ID. If both email and ID are provided, ID will take priority.
   */
  async inviteUserToTeam(
    request: InviteUserToTeamRequest,
    options?: RequestOptions,
  ): Promise<InviteUserToTeamResponseBody> {
    return unwrapAsync(teamsInviteUserToTeam(
      this,
      request,
      options,
    ));
  }

  /**
   * Request access to a team
   *
   * @remarks
   * Request access to a team as a member. An owner has to approve the request. Only 10 users can request access to a team at the same time.
   */
  async requestAccessToTeam(
    request: RequestAccessToTeamRequest,
    options?: RequestOptions,
  ): Promise<RequestAccessToTeamResponseBody> {
    return unwrapAsync(teamsRequestAccessToTeam(
      this,
      request,
      options,
    ));
  }

  /**
   * Get access request status
   *
   * @remarks
   * Check the status of a join request. It'll respond with a 404 if the request has been declined. If no `userId` path segment was provided, this endpoint will instead return the status of the authenticated user.
   */
  async getTeamAccessRequest(
    request: GetTeamAccessRequestRequest,
    options?: RequestOptions,
  ): Promise<GetTeamAccessRequestResponseBody> {
    return unwrapAsync(teamsGetTeamAccessRequest(
      this,
      request,
      options,
    ));
  }

  /**
   * Join a team
   *
   * @remarks
   * Join a team with a provided invite code or team ID.
   */
  async joinTeam(
    request: JoinTeamRequest,
    options?: RequestOptions,
  ): Promise<JoinTeamResponseBody> {
    return unwrapAsync(teamsJoinTeam(
      this,
      request,
      options,
    ));
  }

  /**
   * Update a Team Member
   *
   * @remarks
   * Update the membership of a Team Member on the Team specified by `teamId`, such as changing the _role_ of the member, or confirming a request to join the Team for an unconfirmed member. The authenticated user must be an `OWNER` of the Team.
   */
  async updateTeamMember(
    request: UpdateTeamMemberRequest,
    options?: RequestOptions,
  ): Promise<UpdateTeamMemberResponseBody> {
    return unwrapAsync(teamsUpdateTeamMember(
      this,
      request,
      options,
    ));
  }

  /**
   * Remove a Team Member
   *
   * @remarks
   * Remove a Team Member from the Team, or dismiss a user that requested access, or leave a team.
   */
  async removeTeamMember(
    request: RemoveTeamMemberRequest,
    options?: RequestOptions,
  ): Promise<RemoveTeamMemberResponseBody> {
    return unwrapAsync(teamsRemoveTeamMember(
      this,
      request,
      options,
    ));
  }

  /**
   * Get a Team
   *
   * @remarks
   * Get information for the Team specified by the `teamId` parameter.
   */
  async getTeam(
    request: GetTeamRequest,
    options?: RequestOptions,
  ): Promise<Team> {
    return unwrapAsync(teamsGetTeam(
      this,
      request,
      options,
    ));
  }

  /**
   * Update a Team
   *
   * @remarks
   * Update the information of a Team specified by the `teamId` parameter. The request body should contain the information that will be updated on the Team.
   */
  async patchTeam(
    request: PatchTeamRequest,
    options?: RequestOptions,
  ): Promise<Team> {
    return unwrapAsync(teamsPatchTeam(
      this,
      request,
      options,
    ));
  }

  /**
   * List all teams
   *
   * @remarks
   * Get a paginated list of all the Teams the authenticated User is a member of.
   */
  async getTeams(
    request: GetTeamsRequest,
    options?: RequestOptions,
  ): Promise<GetTeamsResponseBody> {
    return unwrapAsync(teamsGetTeams(
      this,
      request,
      options,
    ));
  }

  /**
   * Create a Team
   *
   * @remarks
   * Create a new Team under your account. You need to send a POST request with the desired Team slug, and optionally the Team name.
   */
  async createTeam(
    request: CreateTeamRequestBody,
    options?: RequestOptions,
  ): Promise<CreateTeamResponseBody> {
    return unwrapAsync(teamsCreateTeam(
      this,
      request,
      options,
    ));
  }

  /**
   * Delete a Team
   *
   * @remarks
   * Delete a team under your account. You need to send a `DELETE` request with the desired team `id`. An optional array of reasons for deletion may also be sent.
   */
  async deleteTeam(
    request: DeleteTeamRequest,
    options?: RequestOptions,
  ): Promise<DeleteTeamResponseBody> {
    return unwrapAsync(teamsDeleteTeam(
      this,
      request,
      options,
    ));
  }

  /**
   * Delete a Team invite code
   *
   * @remarks
   * Delete an active Team invite code.
   */
  async deleteTeamInviteCode(
    request: DeleteTeamInviteCodeRequest,
    options?: RequestOptions,
  ): Promise<DeleteTeamInviteCodeResponseBody> {
    return unwrapAsync(teamsDeleteTeamInviteCode(
      this,
      request,
      options,
    ));
  }
}
